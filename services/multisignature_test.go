package services

import (
	"crypto-exercise-api-danny/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	WantRedeemScript = "522103ad7d69714a5ecdb89aa1e367986b216d75d7186599445bbaa1f390ff1aa0d0bb21037a1630e249ea78f5ecc0d1f99234c4fc26d1c41d7c1e613453678b9d2f0e208921033aff88a08facd0c5234b0e90fd48f6f02f520c5c191f4643fada6d0e8ccc355153ae"
	WantP2shAddress  = "34JZgPaqAhb4a7HJ71HNS1y8fu7qDKfVSb"
)

func TestGenMultisignature(t *testing.T) {
	service := &MultisignatureService{}
	args := models.MultisignatureBody{
		NumberOfApprove: 2,
		ParticipantPublicKeys: []string{
			WantAddressList[7].PublicKey,
			WantAddressList[8].PublicKey,
			WantAddressList[9].PublicKey,
		},
	}
	res, err := service.GenMultisignature(args)

	assert.Equal(t, err, nil)
	if assert.NotNil(t, res) {
		assert.Equal(t, res.RedeemScript, WantRedeemScript)
		assert.Equal(t, res.P2shAddress, WantP2shAddress)
	}
}
