package services

import (
	"bytes"
	"crypto-exercise-api-danny/models"
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
	"strconv"
	"strings"
)

type DerivationService struct{}

func (s *DerivationService) GenDerivationPath(args models.DerivationPathBody) (*models.DerivationPathResp, error) {
	bip32RootKey, err := hdkeychain.NewKeyFromString(args.Bip32RootKeyStr)
	if err != nil {
		return nil, err
	}

	bip32DerivationPath := getBip32DerivationPath(args)
	bip44DerivationPath := getBip44DerivationPath(args)

	var bip32ExtendedKey *hdkeychain.ExtendedKey
	var bip44ExtendedKey *hdkeychain.ExtendedKey
	var bip32ExtendedPublicKey *hdkeychain.ExtendedKey
	var bip44ExtendedPublicKey *hdkeychain.ExtendedKey
	bip32ExtendedKey, err = genExtendedKey(bip32DerivationPath, bip32RootKey)
	if err != nil {
		return nil, err
	}
	bip44ExtendedKey, err = genExtendedKey(bip44DerivationPath, bip32RootKey)
	if err != nil {
		return nil, err
	}
	bip32ExtendedPublicKey, err = bip32ExtendedKey.Neuter()
	if err != nil {
		return nil, err
	}
	bip44ExtendedPublicKey, err = bip44ExtendedKey.Neuter()
	if err != nil {
		return nil, err
	}

	return &models.DerivationPathResp{
		BipPathLevel:      args.BipPathLevel,
		AccountPrivateKey: bip44ExtendedKey.String(),
		AccountPublicKey:  bip44ExtendedPublicKey.String(),
		Bip32PrivateKey:   bip32ExtendedKey.String(),
		Bip32PublicKey:    bip32ExtendedPublicKey.String(),
		DerivationPath:    bip32DerivationPath,
	}, nil
}

func (s *DerivationService) GenDerivationAddresses(args models.DerivationAddressBody) ([]*models.DerivationAddress, error) {
	res := make([]*models.DerivationAddress, args.PageSize)
	for i := 0; i < args.PageSize; i++ {
		var buffer bytes.Buffer
		var extendedKey *hdkeychain.ExtendedKey
		var addr *btcutil.AddressPubKeyHash
		var privateKey *btcec.PrivateKey
		var publicKey *btcec.PublicKey
		var privateKeyWIF *btcutil.WIF

		buffer.WriteString(args.Bip32DerivationPath)
		buffer.WriteString("/")
		buffer.WriteString(strconv.Itoa(args.StartIdx))
		addrDerivationPath := buffer.String()
		bip32MasterKey, err := hdkeychain.NewKeyFromString(args.Bip32RootKeyStr)
		if err != nil {
			return nil, err
		}
		extendedKey, err = genExtendedKey(addrDerivationPath, bip32MasterKey)
		if err != nil {
			return nil, err
		}
		addr, err = extendedKey.Address(&chaincfg.MainNetParams)
		if err != nil {
			return nil, err
		}
		publicKey, err = extendedKey.ECPubKey()
		if err != nil {
			return nil, err
		}
		privateKey, err = extendedKey.ECPrivKey()
		if err != nil {
			return nil, err
		}
		privateKeyWIF, err = btcutil.NewWIF(privateKey, &chaincfg.MainNetParams, true)
		if err != nil {
			return nil, err
		}

		res[i] = &models.DerivationAddress{
			Id:         args.StartIdx,
			Path:       addrDerivationPath,
			Address:    addr.String(),
			PrivateKey: privateKeyWIF.String(),
			PublicKey:  hex.EncodeToString(publicKey.SerializeCompressed()),
		}
		args.StartIdx++
	}

	return res, nil
}

func getBip32DerivationPath(args models.DerivationPathBody) string {
	var buffer bytes.Buffer

	buffer.WriteString("m/")
	buffer.WriteString(strconv.Itoa(args.Purpose))
	buffer.WriteString("'/")
	buffer.WriteString(strconv.Itoa(args.Coin))
	buffer.WriteString("'/")
	buffer.WriteString(strconv.Itoa(args.Account))
	buffer.WriteString("'/")
	buffer.WriteString(strconv.Itoa(args.Change))

	return buffer.String()
}

func getBip44DerivationPath(args models.DerivationPathBody) string {
	var buffer bytes.Buffer

	buffer.WriteString("m/")
	buffer.WriteString(strconv.Itoa(args.Purpose))
	buffer.WriteString("'/")
	buffer.WriteString(strconv.Itoa(args.Coin))
	buffer.WriteString("'/")
	buffer.WriteString(strconv.Itoa(args.Account))
	buffer.WriteString("'")

	return buffer.String()
}

func genExtendedKey(path string, bip32MasterKey *hdkeychain.ExtendedKey) (*hdkeychain.ExtendedKey, error) {
	extendedKey := bip32MasterKey
	pathSlice := strings.Split(path, "/")
	for _, child := range pathSlice {
		if child == "m" {
			continue
		}
		hardened := strings.Contains(child, "'")
		asNormalChild, err := strconv.ParseUint(strings.Replace(child, "'", "", 1), 10, 32)
		if err != nil {
			return nil, err
		}
		childIdx := uint32(asNormalChild)
		if hardened {
			childIdx = hdkeychain.HardenedKeyStart + childIdx
		}
		extendedKey, err = extendedKey.Derive(childIdx)
	}

	return extendedKey, nil
}
