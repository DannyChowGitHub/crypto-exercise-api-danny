package controllers

import (
	"crypto-exercise-api-danny/consts"
	"crypto-exercise-api-danny/models"
	"crypto-exercise-api-danny/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MnemonicController struct {
	service *services.MnemonicService
}

func NewMnemonic() *MnemonicController {
	return &MnemonicController{service: &services.MnemonicService{}}
}

func (c MnemonicController) GenMnemonic(ctx *gin.Context) {
	var body models.MnemonicBody
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var res *models.MnemonicResp
	res, err = c.service.GenMnemonic(body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, consts.InternalErrorMsg)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
