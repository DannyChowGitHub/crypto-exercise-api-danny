package routes

import (
	"crypto-exercise-api-danny/controllers"
	"github.com/gin-gonic/gin"
)

func Multisignature(engine *gin.Engine) {
	ctrl := controllers.NewMultisignature()
	group := engine.Group("/multisignature")
	{
		group.POST("", ctrl.GenMultisignature)
	}
}
