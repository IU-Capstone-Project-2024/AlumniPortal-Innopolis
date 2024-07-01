package middleware

import (
	"alumniportal.com/shared/initializers"
	"alumniportal.com/shared/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

func RequireStudent(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		logrus.WithError(err).Warn("Failed to retrieve Authorization cookie")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if token == nil {
		logrus.Warn("Invalid token")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if err != nil {
		logrus.WithError(err).Warn("Failed to parse JWT token")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) >= claims["exp"].(float64) {
			logrus.Warn("Expired token")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var user models.User

		if err := initializers.DB.First(&user, claims["sub"]).Error; err != nil {
			logrus.WithField("user_id", claims["sub"]).WithError(err).Warn("User not found in database")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", user)
	}

	user, exists := c.Get("user")
	if !exists {
		logrus.Warn("User not found in context")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	authenticatedUser := user.(models.User)
	if !authenticatedUser.Verified {
		logrus.Warn("User is not verified")
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	if authenticatedUser.IsAlumni {
		logrus.Warn("User is alumni")
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	if authenticatedUser.IsAdmin {
		logrus.Warn("User is admin")
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id": authenticatedUser.ID,
		"email":   authenticatedUser.Email,
	}).Info("User authenticated as student")

	c.Next()
}
