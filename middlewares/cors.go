package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var origins = []string{"http://localhost:3000"}
var headers = []string{
	"Content-Type", "Content-Length", "Accept-Encoding",
	"X-CSRF-Token", "Authorization", "accept", "origin",
	"Cache-Control", "X-Requested-With",
}

func Cross() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: origins,
		AllowHeaders: headers,
	})
}
