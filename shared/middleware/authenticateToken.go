package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func AuthenticateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Token")
		expectedToken := "Bearer " + os.Getenv("API_TOKEN")
		if token != expectedToken {
			logrus.Warnf("Unauthorized API access attempt from %s", c.ClientIP())
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}
		c.Next()
	}
}
