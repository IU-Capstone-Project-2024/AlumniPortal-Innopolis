package middleware

import (
	"AlumniPortal-Innopolis/shared/initializers"
	"AlumniPortal-Innopolis/shared/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var body struct {
	Name     string
	LastName string
	Surname  string
	Email    string
	Password string
}

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
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
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) >= claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var user models.User

		initializers.DB.First(&user, claims["sub"])
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", user)

	}

	c.Next()
}
