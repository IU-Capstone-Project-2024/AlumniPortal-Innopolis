package middleware

import (
	"AlumniPortal-Innopolis/shared/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequireAdminRights(c *gin.Context) {
	RequireAuth(c)
	RequireVerify(c)

	user, exists := c.Get("user")
	if !exists {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	authenticatedUser := user.(models.User)
	if !authenticatedUser.IsAdmin {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	c.Next()
}
