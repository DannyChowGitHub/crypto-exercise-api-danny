package routes

import (
	"bip-api/controllers"
	"github.com/gin-gonic/gin"
)

func Multisignature(engine *gin.Engine) {
	ctrl := controllers.NewMultisignature()
	group := engine.Group("/multisignature")
	{
		group.POST("", ctrl.GenMultisignature)
	}
}
