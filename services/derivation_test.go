package services

import (
	"crypto-exercise-api-danny/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	ValidRootKeyString        = "xprv9s21ZrQH143K4XQLiQPfrvLJQp8nRxFwkbX6tYaCHE6Kd8U3iGEngvyPhHwXH4ysFbkb7TZ9ifXe8LPbG7NRAtKHTEfvBTRmaK55SomxY6R"
	ValidBip32DerivationOPath = "m/44'/0'/0'/0"
)

const (
	WantAccountPrivateKey = "xprv9yTpmCEtBaBE6UU65PxxioizcNQBegRXKjqMHcHcZzCQLrrnA53y4RwtytUyaRAEmNBBWsGqPQX6wdu3xW4fTieBPifPcEaSuo82k6UtxH9"
	WantAccountPublicKey  = "xpub6CTBAhmn1wjXJxYZBRVy5wfjAQEg499Ngxkx5zhE8KjPDfBvhcNDcEGNq9nhoSBhfrzBQxeeB8wowRs2zhp4dXujjdhrf4gkwJujDNqL2L3"
	WantBip32PrivateKey   = "xprvA1SxXRUpUqMzeovHgyFt55dakocNBUF7WokZhmtMWTkU4Y6TRKq9mkYg7EhAftj3KZxzkfFFF68AhSgnB17DGrcKgvYZsiuujpdyPewnYqX"
	WantBip32PublicKey    = "xpub6ESJvw1iKCvHsHzknzntSDaKJqSravxxt2gAWAHy4oHSwLRbxs9QKYs9xVMqUt39LY2qRSZkzzCpBpcrQp4keUANWVkj2sXbxeR9EvzZJJ8"
	WantDerivationPath    = "m/44'/0'/0'/0"
)

var (
	WantAddressList = [10]models.DerivationAddress{
		{
			Id:         0,
			Path:       "m/44'/0'/0'/0/0",
			Address:    "1B3bonKcGPKfRMjquJ99cEYnj6iqUEizZL",
			PrivateKey: "KyPwTJR4BFoaTWMLsRAe9VV53jDmtD65w7qZRKJUdKJrCJWABTqS",
			PublicKey:  "0210304335e49ed993920235ad7aaf6e0f36ded1c6d21813047fb378626dc70bc8",
		},
		{
			Id:         1,
			Path:       "m/44'/0'/0'/0/1",
			Address:    "1CYHXLHeLMUg2LP4TXq93cjqufiCympiu7",
			PrivateKey: "KyPcerLeVHQo7TqPBAiJjzqpeT6hASszodLankJcnXmUqKrPVKvY",
			PublicKey:  "031d629cdee5281ccba1546af591da9405b11109a631338d7e88fe2fadbc348ba3",
		},
		{
			Id:         2,
			Path:       "m/44'/0'/0'/0/2",
			Address:    "12LX7YHrWsGvBUM49BxDR35QGpoPDbyZU9",
			PrivateKey: "L55uAPrTFWYhK44N1BVtfXQyCHX3p2Sarr1ayzu2bXf5GYahhDHg",
			PublicKey:  "026796c5e243a5f6d01789b664836cb171e53e163e590d62068179c242553755df",
		},
		{
			Id:         3,
			Path:       "m/44'/0'/0'/0/3",
			Address:    "1FLKZ8qPHaxABnW3j5xGkzn32FxHoXTGd6",
			PrivateKey: "KxhDGX8QD5zuD8tmBRXq8MQpapPkBwBoowYKLkdyMFoubXfWHvCA",
			PublicKey:  "03a6b6e2948c20404e582f645cd149b7541a8218b89611d02d35c62401bff3a928",
		},
		{
			Id:         4,
			Path:       "m/44'/0'/0'/0/4",
			Address:    "1AymKBR5fESniPLgP7mAcfAS5cgRoUrmKR",
			PrivateKey: "L3kDkbHvpCTA1hQQz14KX2r7e8Rk4SPYT4uSSzPCoNanEocs9zkU",
			PublicKey:  "023f25bc7c277b063dd65d3cd92a9e84ecd43b807e711a0504f8c8f9c7d419758b",
		},
		{
			Id:         5,
			Path:       "m/44'/0'/0'/0/5",
			Address:    "1G6hdkQqBDfsaYhKDRoyRS67Uz4ZKNVBaS",
			PrivateKey: "L1mhiGySj82saJyG3VBgfmnPXAA6HdzCdgyB3GqzHP84WJbpowzp",
			PublicKey:  "03cbd000fbb9a81584df0e102bbbf6630e1d51d1aba49f55e445c7bc26ef0b74ad",
		},
		{
			Id:         6,
			Path:       "m/44'/0'/0'/0/6",
			Address:    "13UexHGVrmiLQLhEVP6DP6uivg9mib95P6",
			PrivateKey: "L3q7DQf1DtHdeR57wDqzgDLQ8vweBqHggZXUEpWNEVKVchixALdK",
			PublicKey:  "0270485e77c23f8e09c42e35e7b8042de3ac8e8760709479539101c3cd669471de",
		},
		{
			Id:         7,
			Path:       "m/44'/0'/0'/0/7",
			Address:    "1FhgUNpWNLSo9fJS16M86SpuKRddoRx16o",
			PrivateKey: "KwxTfC2kFHvKiXJ8i8MGKjvYCvojTjM8WuoyE4pFKt7SRPCvThba",
			PublicKey:  "03ad7d69714a5ecdb89aa1e367986b216d75d7186599445bbaa1f390ff1aa0d0bb",
		},
		{
			Id:         8,
			Path:       "m/44'/0'/0'/0/8",
			Address:    "1B7TMGG1mJ8ixXnbkpGtu83mBmk38Ryd9P",
			PrivateKey: "Kzxq9vMRX5tLruKfcvpPRTnFzm5s5DDTqUDAYRnViJU7z9ow3m2o",
			PublicKey:  "037a1630e249ea78f5ecc0d1f99234c4fc26d1c41d7c1e613453678b9d2f0e2089",
		},
		{
			Id:         9,
			Path:       "m/44'/0'/0'/0/9",
			Address:    "1P4yS4Rhu2wBQvJeNCLgB5yHV6ArLgbYMs",
			PrivateKey: "KyDvPgoexR7PNtHuzCTTkcyBhjo3H3QgNpEssx5q5WRsQpUpMhwN",
			PublicKey:  "033aff88a08facd0c5234b0e90fd48f6f02f520c5c191f4643fada6d0e8ccc3551",
		},
	}
)

func TestGenDerivationPathCorrectly(t *testing.T) {
	service := DerivationService{}
	args := models.DerivationPathBody{
		BipPathLevel: models.BipPathLevel{
			Purpose: 44,
			Coin:    0,
			Account: 0,
			Change:  0,
		},
		Bip32RootKeyStr: ValidRootKeyString,
	}

	res, err := service.GenDerivationPath(args)
	assert.Equal(t, err, nil)
	if assert.NotNil(t, res) {
		assert.Equal(t, res.DerivationPath, WantDerivationPath)
		assert.Equal(t, res.AccountPrivateKey, WantAccountPrivateKey)
		assert.Equal(t, res.AccountPublicKey, WantAccountPublicKey)
		assert.Equal(t, res.Bip32PrivateKey, WantBip32PrivateKey)
		assert.Equal(t, res.Bip32PublicKey, WantBip32PublicKey)
	}
}

func TestGenDerivationAddresses(t *testing.T) {
	service := DerivationService{}
	args := models.DerivationAddressBody{
		StartIdx:            0,
		PageSize:            10,
		Bip32DerivationPath: ValidBip32DerivationOPath,
		Bip32RootKeyStr:     ValidRootKeyString,
	}

	res, err := service.GenDerivationAddresses(args)
	assert.Equal(t, err, nil)
	if assert.NotNil(t, res) {
		assert.Equal(t, len(res), len(WantAddressList))
		for i, item := range res {
			assert.Equal(t, item.Id, WantAddressList[i].Id)
			assert.Equal(t, item.Address, WantAddressList[i].Address)
			assert.Equal(t, item.PublicKey, WantAddressList[i].PublicKey)
			assert.Equal(t, item.Path, WantAddressList[i].Path)
			assert.Equal(t, item.PrivateKey, WantAddressList[i].PrivateKey)
		}
	}
}
