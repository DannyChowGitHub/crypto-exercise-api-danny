package main

import (
	"crypto-exercise-api-danny/libs"
	"crypto-exercise-api-danny/middlewares"
	"crypto-exercise-api-danny/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	defer libs.CatchPanic()

	router := gin.Default()
	router.Use(middlewares.Cross())
	router.SetTrustedProxies(nil)
	routes.Mnemonic(router)
	routes.Multisignature(router)
	routes.DerivationPath(router)

	router.Run("localhost:8095")
}
