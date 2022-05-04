package models

type DerivationPathResp struct {
	BipPathLevel
	AccountPrivateKey string `json:"accountPrivateKey" form:"accountPrivateKey"`
	AccountPublicKey  string `json:"accountPublicKey" form:"accountPublicKey"`
	Bip32PrivateKey   string `json:"bip32PrivateKey" form:"bip32PrivateKey"`
	Bip32PublicKey    string `json:"bip32PublicKey" form:"bip32PublicKey"`
	DerivationPath    string `json:"derivationPath" form:"derivationPath"`
}

type BipPathLevel struct {
	Purpose int `json:"purpose" form:"purpose"`
	Coin    int `json:"coin" form:"coin"`
	Account int `json:"account" form:"account"`
	Change  int `json:"change" form:"change"`
}

type DerivationPathBody struct {
	BipPathLevel
	Bip32RootKeyStr string `json:"bip32RootKeyStr" form:"bip32RootKeyStr" binding:"required"`
}

type DerivationAddress struct {
	Id         int    `json:"id"`
	Path       string `json:"path"`
	Address    string `json:"address"`
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}

type DerivationAddressBody struct {
	StartIdx            int    `json:"startIdx" form:"startIdx"`
	PageSize            int    `json:"pageSize" form:"pageSize" binding:"required,min=10,max=100"`
	Bip32DerivationPath string `json:"bip32DerivationPath" form:"bip32DerivationPath" binding:"required"`
	Bip32RootKeyStr     string `json:"bip32RootKeyStr" form:"bip32RootKeyStr" binding:"required"`
}
