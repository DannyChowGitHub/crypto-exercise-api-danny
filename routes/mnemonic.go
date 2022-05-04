package routes

import (
	"bip-api/controllers"
	"github.com/gin-gonic/gin"
)

func Mnemonic(engine *gin.Engine) {
	ctrl := controllers.NewMnemonic()
	group := engine.Group("/mnemonic")
	{
		group.POST("", ctrl.GenMnemonic)
	}
}
