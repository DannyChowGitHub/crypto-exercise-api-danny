package test

import (
	"crypto-exercise-api-danny/models"
	"crypto-exercise-api-danny/services"
	"testing"
)

func TestGenMnemonicLenInvalid(t *testing.T) {
	service := services.MnemonicService{}
	args := models.MnemonicBody{
		NumOfWords: 3,
		Lang:       "english",
		Password:   "",
	}

	_, err := service.GenMnemonic(args)
	if err != nil {
		t.Errorf("Error occur with the args%v", args)
	}
}
