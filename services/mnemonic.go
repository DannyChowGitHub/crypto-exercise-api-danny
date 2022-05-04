package services

import (
	"bip-api/libs"
	"bip-api/models"
	"encoding/hex"
	"errors"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
)

type MnemonicService struct{}

var (
	ErrLengthOfWordsInvalid = errors.New("invalid length of words")
)

func (s *MnemonicService) GenMnemonic(args models.MnemonicBody) (res *models.MnemonicResp, err error) {
	libs.SetWordList(args.Lang)

	var entropy []byte
	entropy, err = newEntropy(args.NumOfWords)
	if err != nil {
		return res, err
	}

	var mnemonic string
	mnemonic, err = bip39.NewMnemonic(entropy)
	if err != nil {
		return res, err
	}

	seed := bip39.NewSeed(mnemonic, args.Password)
	var masterKey *hdkeychain.ExtendedKey
	masterKey, err = hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return res, err
	}

	res = &models.MnemonicResp{
		Words:   mnemonic,
		Seed:    hex.EncodeToString(seed),
		RootKey: masterKey.String(),
	}
	return res, err
}

func newEntropy(numOfWords int) ([]byte, error) {
	if numOfWords < 12 && numOfWords%3 != 0 {
		return nil, ErrLengthOfWordsInvalid
	}

	bitsSize := numOfWords * 11 * 32 / 33
	entropy, err := bip39.NewEntropy(bitsSize)

	return entropy, err
}
