package services

import (
	"crypto-exercise-api-danny/models"
	"encoding/hex"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
)

type MultisignatureService struct{}

func (s *MultisignatureService) GenMultisignature(args models.MultisignatureBody) (*models.MultisignatureResp, error) {
	publicKeys := make([]*btcutil.AddressPubKey, len(args.ParticipantPublicKeys))
	for i, keyStr := range args.ParticipantPublicKeys {
		compressedPubKey, err := hex.DecodeString(keyStr)
		var pubKey *btcutil.AddressPubKey
		pubKey, err = btcutil.NewAddressPubKey(compressedPubKey, &chaincfg.MainNetParams)
		if err != nil {
			return nil, err
		}
		publicKeys[i] = pubKey
	}

	p2ms, err := txscript.MultiSigScript(publicKeys, args.NumberOfApprove)
	if err != nil {
		return nil, err
	}
	var addr *btcutil.AddressScriptHash
	addr, err = btcutil.NewAddressScriptHash(p2ms, &chaincfg.MainNetParams)

	return &models.MultisignatureResp{
		RedeemScript: hex.EncodeToString(p2ms),
		P2shAddress:  addr.String(),
	}, nil
}
