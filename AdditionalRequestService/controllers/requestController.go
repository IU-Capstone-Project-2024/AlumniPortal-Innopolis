package controllers

import (
	"AdditionalRequestService/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"shared/initializers"
	sharedModels "shared/models"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startDate, err := time.Parse(time.RFC3339, input.PassStartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
		return
	}

	expirationDate, err := time.Parse(time.RFC3339, input.PassExpirationDate)
	if err != nil {
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, passRequest)
}

func DeletePassRequest(c *gin.Context) {
	var passRequest models.PassRequest
	if err := initializers.DB.Where("id = ? AND status = ?", c.Param("id"), models.Unverified).First(&passRequest).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pass request not found or already accepted/declined"})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	if passRequest.UserID != user.(sharedModels.User).ID && user.(sharedModels.User).IsAdmin == false {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized: User is not the owner of this pass request"})
		return
	}

	initializers.DB.Delete(&passRequest)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func UpdatePassRequest(c *gin.Context) {
	var passRequest models.PassRequest
	if err := initializers.DB.Where("id = ? AND status = ?", c.Param("id"), models.Unverified).First(&passRequest).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pass request not found or already accepted/declined"})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	if passRequest.UserID != user.(sharedModels.User).ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized: User is not the owner of this pass request"})
		return
	}

	var input PassRequestInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startDate, err := time.Parse(time.RFC3339, input.PassStartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
		return
	}

	expirationDate, err := time.Parse(time.RFC3339, input.PassExpirationDate)
	if err != nil {
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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
		c.JSON(http.StatusNotFound, gin.H{"error": "Pass request not found"})
		return
	}

	passRequest.Status = status
	if err := initializers.DB.Save(&passRequest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, passRequest)
}

func GetCurrentUserRequests(c *gin.Context) {
	user, _ := c.Get("user")
	var passRequests []models.PassRequest
	if err := initializers.DB.Where("user_id = ?", user.(sharedModels.User).ID).Find(&passRequests).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, passRequests)
}

func GetUnverifiedRequests(c *gin.Context) {
	var passRequests []models.PassRequest
	if err := initializers.DB.Where("status = ?", models.Unverified).Find(&passRequests).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, passRequests)
}
