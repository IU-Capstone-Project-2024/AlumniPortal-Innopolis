package controllers

import (
	"alumniportal.com/shared/helpers"
	"net/http"

	"alumniportal.com/shared/initializers"
	sharedModels "alumniportal.com/shared/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateVolunteerRequest(c *gin.Context) {

	user, _ := c.Get("user")
	project, _ := c.Get("project")
	passRequest := sharedModels.Volunteer{
		UserID:    user.(sharedModels.User).ID,
		ProjectID: user.(sharedModels.Project).ID,
		Status:    sharedModels.UnverifiedVolunteer,
	}
	if err := initializers.DB.Create(&passRequest).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to create pass request")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id":          user.(sharedModels.User).ID,
		"project_id":       project.(sharedModels.Project).ID,
		"volunteer_req_id": passRequest.ID,
	}).Info("Pass request created successfully")
	c.JSON(http.StatusOK, passRequest)

}
func DeleteVolunteerRequest(c *gin.Context) {
	var volunteerModel sharedModels.Volunteer
	if err := initializers.DB.Where("id = ? AND status = ?", c.Param("id"), sharedModels.UnverifiedVolunteer).First(&volunteerModel).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"volunteer_req_id": c.Param("id"),
			"error":            err.Error(),
		}).Error("Pass request not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Pass request not found"})
		return
	}
	user, exists := c.Get("user")
	if !exists {
		logrus.Warn("User not authenticated for DeletePassRequest")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	if volunteerModel.UserID != user.(sharedModels.User).ID && user.(sharedModels.User).Role != helpers.Admin {
		logrus.WithFields(logrus.Fields{
			"user_id":          user.(sharedModels.User).ID,
			"volunteer_req_id": volunteerModel.ID,
		}).Warn("Unauthorized user attempting to delete pass request")
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized: User is not the owner of this pass request"})
		return
	}

	initializers.DB.Delete(&volunteerModel)
	logrus.WithFields(logrus.Fields{
		"user_id":          user.(sharedModels.User).ID,
		"volunteer_req_id": volunteerModel.ID,
	}).Info("Pass request deleted successfully")
	c.JSON(http.StatusOK, gin.H{"data": true})

}

func updateVolunteerRequestStatus(c *gin.Context, status sharedModels.VolunteerVerificationStatus) {
	var volunteerModel sharedModels.Volunteer
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&volunteerModel).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"volunteer_req_id": c.Param("id"),
			"error":            err.Error(),
		}).Error("Pass request not found for updatePassRequestStatus")
		c.JSON(http.StatusNotFound, gin.H{"error": "Pass request not found"})
		return
	}

	volunteerModel.Status = status
	if err := initializers.DB.Save(&volunteerModel).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"volunteer_req_id": volunteerModel.ID,
			"error":            err.Error(),
		}).Error("Failed to update pass request status")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"volunteer_req_id": volunteerModel.ID,
		"status":           status,
	}).Info("Pass request status updated successfully")
	c.JSON(http.StatusOK, volunteerModel)
}
func AcceptVolunteerRequest(c *gin.Context) {
	updateVolunteerRequestStatus(c, sharedModels.AcceptedVolunteer)
}
func DeclineVolunteerRequest(c *gin.Context) {
	updateVolunteerRequestStatus(c, sharedModels.DeclinedVolunteer)
}
func GetStudentVolunteerRequest(c *gin.Context) {
	var volunteerModel sharedModels.Volunteer
	if err := initializers.DB.Where("id = ? AND status = ?", c.Param("id"), sharedModels.UnverifiedVolunteer).Preload("User").First(&volunteerModel).Error; err != nil {
		logrus.WithFields(logrus.Fields{}).Error("Failed to get student volunteer request")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	logrus.WithFields(logrus.Fields{
		"volunteer_req_id": volunteerModel.ID,
	}).Info("Fetched student volunteer request successfully")
	c.JSON(http.StatusOK, volunteerModel)
}
func GetUnverifiedVolunteers(c *gin.Context) {
	var volunteers []sharedModels.Volunteer
	if err := initializers.DB.Where("status = ?", sharedModels.UnverifiedVolunteer).Preload("User").Preload("Project").Find(&volunteers).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to get unverified volunteers")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	logrus.WithFields(logrus.Fields{
		"count": len(volunteers),
	}).Info("Fetched unverified volunteers successfully")
	c.JSON(http.StatusOK, volunteers)
}
func GetCurrentUserRequests(c *gin.Context) {
	user, _ := c.Get("user")
	var volunteerRequests []sharedModels.Volunteer
	if err := initializers.DB.Where("user_id = ?", user.(sharedModels.User).ID).Preload("User").Preload("Project").Find(&volunteerRequests).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"user_id": user.(sharedModels.User).ID,
			"error":   err.Error(),
		}).Error("Failed to get current user pass requests")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id": user.(sharedModels.User).ID,
		"count":   len(volunteerRequests),
	}).Info("Fetched current user pass requests successfully")
	c.JSON(http.StatusOK, volunteerRequests)
}
func GetVolunteerRequest(c *gin.Context) {
	var volunteerRequest sharedModels.Volunteer

	user, _ := c.Get("user")
	project, _ := c.Get("project")

	if err := initializers.DB.Where("id = ? AND user_id = ? AND project_id = ? AND role = ?", c.Param("id"), user.(sharedModels.User).ID, project.(sharedModels.Project).ID, user.(sharedModels.User).Role).Preload("User").Preload("Project").First(&volunteerRequest).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"user_id":          user.(sharedModels.User).ID,
			"volunteer_req_id": c.Param("id"),
			"error":            err.Error(),
		}).Error("Failed to get pass request")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logrus.WithFields(logrus.Fields{
		"user_id":          user.(sharedModels.User).ID,
		"volunteer_req_id": volunteerRequest.ID,
	}).Info("Fetched pass request successfully")
	c.JSON(http.StatusOK, volunteerRequest)
}
