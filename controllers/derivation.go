package controllers

import (
	"crypto-exercise-api-danny/consts"
	"crypto-exercise-api-danny/models"
	"crypto-exercise-api-danny/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DerivationController struct {
	service *services.DerivationService
}

func NewDerivation() *DerivationController {
	return &DerivationController{service: &services.DerivationService{}}
}

func (c DerivationController) GenDerivationPath(ctx *gin.Context) {
	var body models.DerivationPathBody
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var res *models.DerivationPathResp
	res, err = c.service.GenDerivationPath(body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, consts.InternalErrorMsg)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c DerivationController) GenDerivationAddresses(ctx *gin.Context) {
	var body models.DerivationAddressBody
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var res []*models.DerivationAddress
	res, err = c.service.GenDerivationAddresses(body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, consts.InternalErrorMsg)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
