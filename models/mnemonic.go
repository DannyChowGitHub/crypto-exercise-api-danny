package models

type MnemonicResp struct {
	Words   string `json:"words" form:"words"`
	Seed    string `json:"seed" form:"seed"`
	RootKey string `json:"rootKey" form:"rootKey"`
}

type MnemonicBody struct {
	NumOfWords int    `json:"numOfWords" form:"numOfWords" binding:"required"`
	Lang       string `json:"lang" form:"lang"`
	Password   string `json:"password" form:"password"`
}
