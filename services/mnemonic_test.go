package services

import (
	"crypto-exercise-api-danny/models"
	"encoding/hex"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGenMnemonicCorrectly(t *testing.T) {
	service := MnemonicService{}
	args := models.MnemonicBody{
		NumOfWords: 12,
		Lang:       "english",
		Password:   "",
	}
	res, err := service.GenMnemonic(args)
	if err != nil {
		t.Errorf("%v", err)
	}

	assert.Equal(t, err, nil)
	wordList := strings.Split(res.Words, " ")
	assert.Equal(t, len(wordList), args.NumOfWords)

	originSeed, err := hex.DecodeString(res.Seed)
	if err != nil {
		t.Errorf("%v", err)
	}
	masterKey, err := hdkeychain.NewMaster(originSeed, &chaincfg.MainNetParams)
	if err != nil {
		t.Errorf("%v", err)
	}
	assert.Equal(t, res.RootKey, masterKey.String())
}

func TestGenMnemonicInvalidLength(t *testing.T) {
	service := MnemonicService{}
	args := models.MnemonicBody{
		NumOfWords: 3,
		Lang:       "english",
		Password:   "",
	}
	_, err := service.GenMnemonic(args)
	if assert.NotNil(t, err) {
		assert.Equal(t, err.Error(), ErrLengthOfWordsInvalid.Error())
	}
}

func TestGenMnemonicNotTripleOf3(t *testing.T) {
	service := MnemonicService{}
	args := models.MnemonicBody{
		NumOfWords: 13,
		Lang:       "english",
		Password:   "",
	}
	_, err := service.GenMnemonic(args)
	if assert.NotNil(t, err) {
		assert.Equal(t, err.Error(), ErrLengthOfWordsInvalid.Error())
	}
}
