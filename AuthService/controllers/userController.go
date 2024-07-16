package controllers

import (
	"net/http"
	"os"
	"time"

	"alumniportal.com/shared/initializers"
	"alumniportal.com/shared/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var body struct {
	Name      string `json:"name"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	IsAlumni  *bool  `json:"isAlumni"`
	IsStudent *bool  `json:"isAdmin"`
}

type EditUser struct {
	Name             *string `json:"name"`
	LastName         *string `json:"lastName"`
	Email            *string `json:"email"`
	Role             *string `json:"role"`
	Specialization   *string `json:"specialization"`
	PortfolioLink    *string `json:"portfolioLink"`
	SocialsLink      *string `json:"socialsLink"`
	AvailableCustdev *bool   `gorm:"availableCustdev"`
	Password         *string `json:"password"`
}

func Signup(c *gin.Context) {
	if c.ShouldBindJSON(&body) != nil {
		logrus.WithFields(logrus.Fields{
			"error": "Failed to read body",
		}).Error("Signup error")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body.",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to hash password")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	if body.IsAlumni == nil || body.IsStudent == nil || (*body.IsAlumni && *body.IsStudent) || (!*body.IsAlumni && !*body.IsStudent) {
		logrus.Warn("Invalid alumni/student status")
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid alumni/student status"})
		return
	}
	user := models.User{
		Name:     body.Name,
		LastName: body.LastName,
		Email:    body.Email,
		Password: string(hash),
	}

	if *body.IsAlumni {
		user.IsAlumni = true
	} else {
		user.IsStudent = true
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		logrus.WithFields(logrus.Fields{
			"error": result.Error.Error(),
		}).Error("Failed to create user")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id": user.ID,
	}).Info("User created successfully")
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	if c.Bind(&body) != nil {
		logrus.WithFields(logrus.Fields{
			"error": "Failed to read body",
		}).Error("Login error")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body.",
		})
		return
	}

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		logrus.WithFields(logrus.Fields{
			"email": body.Email,
		}).Warn("Invalid email or password")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Email or Password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"email": body.Email,
		}).Warn("Invalid email or password")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Email or Password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 48).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to sign token")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)
	logrus.WithFields(logrus.Fields{
		"user_id": user.ID,
	}).Info("User logged in successfully")
	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	if user == nil {
		logrus.Warn("User not found during validation")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found",
		})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id": user.(models.User).ID,
	}).Info("User validated successfully")
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func GetInfo(c *gin.Context) {
	user, _ := c.Get("user")

	if user == nil {
		logrus.Warn("User not found during info retrieval")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found",
		})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id": user.(models.User).ID,
	}).Info("User info retrieved successfully")
	c.JSON(http.StatusOK, gin.H{
		"Name":           user.(models.User).Name,
		"LastName":       user.(models.User).LastName,
		"Email":          user.(models.User).Email,
		"Role":           user.(models.User).Role,
		"Specialization": user.(models.User).Specialization,
		"Portfolio":      user.(models.User).PortfolioLink,
		"Socials":        user.(models.User).SocialsLink,
		"IsAlumni":       user.(models.User).IsAlumni,
		"IsAdmin":        user.(models.User).IsAdmin,
	})
}
func Edit(c *gin.Context) {
	user, _ := c.Get("user")
	if user == nil {
		logrus.Warn("User not found during info retrieval")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found",
		})
		return
	}
	var input EditUser
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to bind JSON for UpdatePassRequest")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var updateData models.User
	if input.Name != nil {
		updateData.Name = input.Name
	}
	if input.LastName != nil {
		updateData.LastName = input.LastName
	}
	if input.LastName != nil {
		updateData.Email = input.Email
	}
	if input.PortfolioLink != nil {
		updateData.PortfolioLink = input.PortfolioLink
	}
	if input.SocialsLink != nil {
		updateData.SocialsLink = input.SocialsLink
	}
	if input.Role != nil {
		updateData.Role = input.Role
	}
	if input.Specialization != nil {
		updateData.Specialization = input.Specialization
	}
	if input.AvailableCustdev != nil {
		updateData.AvailableCustdev = input.AvailableCustdev
	}
	if input.Password != nil {
		hash, err := bcrypt.GenerateFromPassword([]byte(*input.Password), 10)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("Failed to hash password")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to hash password",
			})
			return
		}
		updateData.Password = string(hash)
	}

	if err := initializers.DB.Model(&user).Updates(updateData).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to update user request")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id": user.(models.User).ID,
	}).Info("User updated successfully")
	c.JSON(http.StatusOK, user)

}
