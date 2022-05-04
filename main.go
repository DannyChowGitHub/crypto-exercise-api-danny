package main

import (
	"bip-api/libs"
	"bip-api/middlewares"
	"bip-api/routes"
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
