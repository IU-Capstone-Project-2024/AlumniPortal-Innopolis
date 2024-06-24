package controllers

import (
	"alumniportal.com/shared/initializers"
	"alumniportal.com/shared/models"
	sharedModels "alumniportal.com/shared/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type PassRequestInput struct {
	PassStartDate      string `json:"pass_start_date" binding:"required"`
	PassExpirationDate string `json:"pass_expiration_date" binding:"required"`
	Message            string `json:"message" binding:"required"`
	PassType           string `json:"pass_type" binding:"required,oneof=Dormitory University"`
}

func CreatePassRequest(c *gin.Context) {
	var input PassRequestInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to bind JSON for CreatePassRequest")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startDate, err := time.Parse(time.RFC3339, input.PassStartDate)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Invalid start date format for CreatePassRequest")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
		return
	}

	expirationDate, err := time.Parse(time.RFC3339, input.PassExpirationDate)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Invalid expiration date format for CreatePassRequest")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expiration date format"})
		return
	}

	user, _ := c.Get("user")
	passRequest := models.PassRequest{
		UserID:             user.(sharedModels.User).ID,
		PassStartDate:      startDate,
		PassExpirationDate: expirationDate,
		Message:            input.Message,
		PassType:           models.PassType(input.PassType),
		Status:             models.Unverified,
	}

	if err := initializers.DB.Create(&passRequest).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to create pass request")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id":    user.(sharedModels.User).ID,
		"request_id": passRequest.ID,
	}).Info("Pass request created successfully")
	c.JSON(http.StatusOK, passRequest)
}

func DeletePassRequest(c *gin.Context) {
	var passRequest models.PassRequest

	user, exists := c.Get("user")

	if !exists {
		logrus.Warn("User not authenticated for DeletePassRequest")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	if user.(sharedModels.User).IsAdmin == true {
		if err := initializers.DB.Where("id = ?", c.Param("id")).First(&passRequest).Error; err != nil {
			logrus.WithFields(logrus.Fields{
				"request_id": c.Param("id"),
				"error":      err.Error(),
			}).Error("Pass request not found for DeletePassRequest")
			c.JSON(http.StatusNotFound, gin.H{"error": "Pass request not found"})
			return
		}
	} else {
		if err := initializers.DB.Where("id = ? AND status = ?", c.Param("id"), models.Unverified).First(&passRequest).Error; err != nil {
			logrus.WithFields(logrus.Fields{
				"request_id": c.Param("id"),
				"error":      err.Error(),
			}).Error("Pass request not found or already accepted/declined for DeletePassRequest")
			c.JSON(http.StatusNotFound, gin.H{"error": "Pass request not found or already accepted/declined"})
			return
		}
	}

	if passRequest.UserID != user.(sharedModels.User).ID && user.(sharedModels.User).IsAdmin == false {
		logrus.WithFields(logrus.Fields{
			"user_id":    user.(sharedModels.User).ID,
			"request_id": passRequest.ID,
		}).Warn("Unauthorized user attempting to delete pass request")
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized: User is not the owner of this pass request"})
		return
	}

	initializers.DB.Delete(&passRequest)
	logrus.WithFields(logrus.Fields{
		"user_id":    user.(sharedModels.User).ID,
		"request_id": passRequest.ID,
	}).Info("Pass request deleted successfully")
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func UpdatePassRequest(c *gin.Context) {
	var passRequest models.PassRequest
	if err := initializers.DB.Where("id = ? AND status = ?", c.Param("id"), models.Unverified).First(&passRequest).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"request_id": c.Param("id"),
			"error":      err.Error(),
		}).Error("Pass request not found or already accepted/declined for UpdatePassRequest")
		c.JSON(http.StatusNotFound, gin.H{"error": "Pass request not found or already accepted/declined"})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		logrus.Warn("User not authenticated for UpdatePassRequest")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	if passRequest.UserID != user.(sharedModels.User).ID {
		logrus.WithFields(logrus.Fields{
			"user_id":    user.(sharedModels.User).ID,
			"request_id": passRequest.ID,
		}).Warn("Unauthorized user attempting to update pass request")
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized: User is not the owner of this pass request"})
		return
	}

	var input PassRequestInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to bind JSON for UpdatePassRequest")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startDate, err := time.Parse(time.RFC3339, input.PassStartDate)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Invalid start date format for UpdatePassRequest")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
		return
	}

	expirationDate, err := time.Parse(time.RFC3339, input.PassExpirationDate)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Invalid expiration date format for UpdatePassRequest")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expiration date format"})
		return
	}

	updateData := models.PassRequest{
		PassStartDate:      startDate,
		PassExpirationDate: expirationDate,
		Message:            input.Message,
		PassType:           models.PassType(input.PassType),
	}

	if err := initializers.DB.Model(&passRequest).Updates(updateData).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to update pass request")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id":    user.(sharedModels.User).ID,
		"request_id": passRequest.ID,
	}).Info("Pass request updated successfully")
	c.JSON(http.StatusOK, passRequest)
}

func ApprovePassRequest(c *gin.Context) {
	updatePassRequestStatus(c, models.Accepted)
}

func DeclinePassRequest(c *gin.Context) {
	updatePassRequestStatus(c, models.Declined)
}

func updatePassRequestStatus(c *gin.Context, status models.PassRequestStatus) {
	var passRequest models.PassRequest
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&passRequest).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"request_id": c.Param("id"),
			"error":      err.Error(),
		}).Error("Pass request not found for updatePassRequestStatus")
		c.JSON(http.StatusNotFound, gin.H{"error": "Pass request not found"})
		return
	}

	passRequest.Status = status
	if err := initializers.DB.Save(&passRequest).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"request_id": passRequest.ID,
			"error":      err.Error(),
		}).Error("Failed to update pass request status")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"request_id": passRequest.ID,
		"status":     status,
	}).Info("Pass request status updated successfully")
	c.JSON(http.StatusOK, passRequest)
}

func GetCurrentUserRequests(c *gin.Context) {
	user, _ := c.Get("user")
	var passRequests []models.PassRequest
	if err := initializers.DB.Where("user_id = ?", user.(sharedModels.User).ID).Preload("User").Find(&passRequests).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"user_id": user.(sharedModels.User).ID,
			"error":   err.Error(),
		}).Error("Failed to get current user pass requests")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id": user.(sharedModels.User).ID,
		"count":   len(passRequests),
	}).Info("Fetched current user pass requests successfully")
	c.JSON(http.StatusOK, passRequests)
}

func GetUnverifiedRequests(c *gin.Context) {
	var passRequests []models.PassRequest

	if err := initializers.DB.Where("status = ?", models.Unverified).Preload("User").Find(&passRequests).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to get unverified pass requests")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"count": len(passRequests),
	}).Info("Fetched unverified pass requests successfully")
	c.JSON(http.StatusOK, passRequests)
}

func GetAdminPassRequest(c *gin.Context) {
	var passRequest models.PassRequest

	if err := initializers.DB.Where("status = ? AND id = ?", models.Unverified, c.Param("id")).Preload("User").First(&passRequest).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"request_id": c.Param("id"),
			"error":      err.Error(),
		}).Error("Failed to get admin pass request")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"request_id": passRequest.ID,
	}).Info("Fetched admin pass request successfully")
	c.JSON(http.StatusOK, passRequest)
}

func GetPassRequest(c *gin.Context) {
	var passRequest models.PassRequest

	user, _ := c.Get("user")

	if err := initializers.DB.Where("id = ? AND user_id = ?", c.Param("id"), user.(sharedModels.User).ID).Preload("User").First(&passRequest).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"user_id":    user.(sharedModels.User).ID,
			"request_id": c.Param("id"),
			"error":      err.Error(),
		}).Error("Failed to get pass request")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id":    user.(sharedModels.User).ID,
		"request_id": passRequest.ID,
	}).Info("Fetched pass request successfully")
	c.JSON(http.StatusOK, passRequest)
}
