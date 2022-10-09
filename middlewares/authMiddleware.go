package middlewares

import (
	"gymdb/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

var AuthMiddleware = gin.BasicAuth(gin.Accounts{
	initializers.AUTH_USERNAME: initializers.AUTH_PASSWORD,
})
