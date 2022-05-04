package controllers

import (
	"bip-api/consts"
	"bip-api/models"
	"bip-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MultisignatureController struct {
	service *services.MultisignatureService
}

func NewMultisignature() *MultisignatureController {
	return &MultisignatureController{service: &services.MultisignatureService{}}
}

func (c MultisignatureController) GenMultisignature(ctx *gin.Context) {
	var body models.MultisignatureBody
	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var res *models.MultisignatureResp
	res, err = c.service.GenMultisignature(body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, consts.InternalErrorMsg)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
