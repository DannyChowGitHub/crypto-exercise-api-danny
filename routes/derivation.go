package routes

import (
	"bip-api/controllers"
	"github.com/gin-gonic/gin"
)

func DerivationPath(engine *gin.Engine) {
	ctrl := controllers.NewDerivation()
	group := engine.Group("/derivation")
	{
		group.POST("/path", ctrl.GenDerivationPath)
		group.POST("/addresses", ctrl.GenDerivationAddresses)
	}
}
