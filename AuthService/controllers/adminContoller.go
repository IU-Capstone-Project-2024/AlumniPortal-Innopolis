package controllers

import (
	"alumniportal.com/shared/initializers"
	"alumniportal.com/shared/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

var userId struct {
	UserID uint `json:"user_id" binding:"required"`
}

func Verify(c *gin.Context) {
	logrus.Info("Verifying user")

	if err := c.ShouldBindJSON(&userId); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to read body")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body. Make sure 'userId' is provided.",
		})
		return
	}

	var user models.User

	if err := initializers.DB.First(&user, userId.UserID).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"user_id": userId.UserID,
			"error":   err.Error(),
		}).Error("User not found")
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	user.Verified = true

	if err := initializers.DB.Save(&user).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"user_id": user.ID,
			"error":   err.Error(),
		}).Error("Failed to update user verification status")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user verification status",
		})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id": user.ID,
	}).Info("User verified successfully")
	c.JSON(http.StatusOK, gin.H{
		"message": "User verified successfully",
	})
}

func DeleteUser(c *gin.Context) {
	logrus.Info("Deleting user")

	if err := c.ShouldBindJSON(&userId); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to read body")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body. Make sure 'userId' is provided.",
		})
		return
	}

	var user models.User

	if err := initializers.DB.First(&user, userId.UserID).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"user_id": userId.UserID,
			"error":   err.Error(),
		}).Error("User not found")
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	if err := initializers.DB.Delete(&user).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"user_id": user.ID,
			"error":   err.Error(),
		}).Error("Failed to delete user")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete user",
		})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id": user.ID,
	}).Info("User deleted successfully")
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
