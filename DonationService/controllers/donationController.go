package controllers

import (
	"net/http"
	"time"

	"alumniportal.com/shared/initializers"
	"alumniportal.com/shared/models"
	sharedModels "alumniportal.com/shared/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/datatypes"
)

type DonationRequestInput struct {
	Amount            float32 `json:"amount" binding:"required"`
	RecurringDonation string  `json:"recurring_donation" binding:"required,oneof=OneTime Monthly Quarterly Yearly"`
	PaymentMethod     string  `json:"payment_method" binding:"required,oneof=CreditCard PayPal BankTransfer"`
}

func CreateDonationRequest(c *gin.Context) {
	var input DonationRequestInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to bind JSON for CreateDonationRequest")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	date := datatypes.Date(time.Now())
	user, _ := c.Get("user")
	project, _ := c.Get("project")
	donationRequest := sharedModels.Donation{
		DonatorID:         user.(sharedModels.User).ID,
		ProjectID:         project.(sharedModels.Project).ID,
		Amount:            input.Amount,
		Date:              date,
		PaymentMethod:     models.PaymentMethod(input.PaymentMethod),
		RecurringDonation: models.RecurringDonation(input.RecurringDonation),
	}
	if err := initializers.DB.Create(&donationRequest).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to create donation request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logrus.WithFields(logrus.Fields{
		"user_id":         user.(sharedModels.User).ID,
		"project_id":      project.(sharedModels.Project).ID,
		"amount":          input.Amount,
		"donation_req_id": donationRequest.ID,
	}).Info("Donation request created successfully")
	c.JSON(http.StatusOK, donationRequest)
}

func GetCurrentUserDonationRequests(c *gin.Context) {
	var donationRequests []sharedModels.Donation
	user, _ := c.Get("user")
	if err := initializers.DB.Where("donator_id = ?", user.(sharedModels.User).ID).Preload("User").Find(&donationRequests).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"user_id": user.(sharedModels.User).ID,
			"error":   err.Error(),
		}).Error("User is not found for GetCurrentUserDonationRequests")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User is not found"})
		return
	}
	logrus.WithFields(logrus.Fields{
		"user_id": user.(sharedModels.User).ID,
		"count":   len(donationRequests),
	}).Info("Fetched current user donation requests successfully")
	c.JSON(http.StatusOK, donationRequests)
}

func GetCurrentProjectDonationRequests(c *gin.Context) {
	var donationRequests []sharedModels.Donation

	projectId := c.Param("id")

	if err := initializers.DB.Where("project_id = ?", projectId).Preload("Project").Preload("User").Find(&donationRequests).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"project_id": projectId,
			"error":      err.Error(),
		}).Error("Project is not found for GetCurrentProjectDonationRequests")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Project is not found"})
		return
	}
	logrus.WithFields(logrus.Fields{
		"project_id": projectId,
		"count":      len(donationRequests),
	}).Info("Fetched current project donation requests successfully")
	c.JSON(http.StatusOK, donationRequests)
}

func GetCurrentDateDonationRequests(c *gin.Context) {
	var donationRequests []sharedModels.Donation
	if err := initializers.DB.Where("date = ?", c.Param("date")).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"date":  c.Param("date"),
			"error": err.Error(),
		}).Error("Date is not found for GetCurrentDatetDonationRequests")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Date is not found"})
		return
	}
	logrus.WithFields(logrus.Fields{
		"date":  c.Param("date"),
		"count": len(donationRequests),
	}).Info("Fetched current date donation requests successfully")
	c.JSON(http.StatusOK, donationRequests)
}
func GetCurrentAmountDonationRequests(c *gin.Context) {
	var donationRequests []sharedModels.Donation
	if err := initializers.DB.Where("amount = ?", c.Param("amount")).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"amount": c.Param("amount"),
			"error":  err.Error(),
		}).Error("Date is not found for GetCurrentAmountDonationRequests")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Date is not found"})
		return
	}
	logrus.WithFields(logrus.Fields{
		"amount": c.Param("amount"),
		"count":  len(donationRequests),
	}).Info("Fetched current amount donation requests successfully")
	c.JSON(http.StatusOK, donationRequests)
}
func GetDonationRequest(c *gin.Context) {
	var donationRequest sharedModels.Donation
	user, _ := c.Get("user")
	project, _ := c.Get("project")
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&donationRequest).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"donation_req_id": donationRequest.ID,
			"error":           err.Error(),
		}).Error("Donation request is not found for GetDonationRequest")
		c.JSON(http.StatusNotFound, gin.H{"error": "Donation request is not found"})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id":         user.(sharedModels.User).ID,
		"project_id":      project.(sharedModels.Project).ID,
		"amount":          donationRequest.Amount,
		"donation_req_id": donationRequest.ID,
	}).Info("Fetched donation successfuly")
	c.JSON(http.StatusOK, donationRequest)
}

func GetAccumulatedSumDonationRequest(c *gin.Context) {
	var donationRequests []sharedModels.Donation
	project, _ := c.Get("project")
	if err := initializers.DB.Where("project_id = ?", project.(sharedModels.Project).ID).Preload("Project").Find(&donationRequests).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"project_id": project.(sharedModels.Project).ID,
			"error":      err.Error(),
		}).Error("Project is not found for GetRemainDonationRequests")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Project is not found"})
		return
	}
	var sum float32 = 0

	for _, value := range donationRequests {
		sum += value.Amount
	}

	logrus.WithFields(logrus.Fields{
		"project_id":  project.(sharedModels.Project).ID,
		"full_amount": sum,
	}).Info("Fetched sum successfuly")
	c.JSON(http.StatusOK, sum)
}
