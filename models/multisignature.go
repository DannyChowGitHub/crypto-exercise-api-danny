package models

type MultisignatureResp struct {
	RedeemScript string `json:"redeemScript"`
	P2shAddress  string `json:"p2shAddress"`
}

type MultisignatureBody struct {
	NumberOfApprove       int      `json:"numberOfApprove" form:"numberOfApprove" binding:"required"`
	ParticipantPublicKeys []string `json:"participantPublicKeys" form:"participantPublicKeys" binding:"required"`
}
