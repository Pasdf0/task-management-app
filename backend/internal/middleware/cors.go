package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors Middleware : Agent CORS Middleware
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()

	config.AllowAllOrigins = true
	config.AllowMethods = append(config.AllowMethods, "GET", "POST", "PUT", "DELETE", "OPTIONS")
	config.AllowHeaders = append(config.AllowHeaders, "Content-Type", "Authorization")

	return cors.New(config)
}
