package controllers

import (
	"alumniportal.com/shared/initializers"
	sharedModels "alumniportal.com/shared/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ParticipantInput struct {
	ProjectID uint
}

func AddParticipant(c *gin.Context) {
	var input ParticipantInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to bind JSON for AddParticipant")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := c.Get("user")
	projectParticipant := sharedModels.ProjectParticipant{
		UserID:    user.(sharedModels.User).ID,
		ProjectID: input.ProjectID,
		Status:    sharedModels.UnverifiedProject,
	}

	if err := initializers.DB.Create(&projectParticipant).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to add participant")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id":    user.(sharedModels.User).ID,
		"project_id": projectParticipant.ProjectID,
		"add_id":     projectParticipant.ID,
	}).Info("Participant added successfully")
	c.JSON(http.StatusOK, projectParticipant)
}

func RemoveParticipant(c *gin.Context) {
	var participant sharedModels.ProjectParticipant

	user, exists := c.Get("user")

	if !exists {
		logrus.Warn("User not authenticated for Removing Participant")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	if user.(sharedModels.User).IsAdmin == true {
		if err := initializers.DB.Where("id = ?", c.Param("id")).First(&participant).Error; err != nil {
			logrus.WithFields(logrus.Fields{
				"participation_id": c.Param("id"),
				"user_id":          participant.UserID,
				"project_id":       participant.ProjectID,
				"error":            err.Error(),
			}).Error("Project not found for Removing participant")
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}
	} else {
		if err := initializers.DB.Where("id = ?", c.Param("id")).First(&participant).Error; err != nil {
			logrus.WithFields(logrus.Fields{
				"participation_id": c.Param("id"),
				"user_id":          participant.UserID,
				"project_id":       participant.ProjectID,
				"error":            err.Error(),
			}).Error("Project not found for Delete participant")
			c.JSON(http.StatusNotFound, gin.H{"error": "Participation not found"})
			return
		}
	}

	if participant.UserID != user.(sharedModels.User).ID && user.(sharedModels.User).IsAdmin == false {
		logrus.WithFields(logrus.Fields{
			"participation_id": participant.ID,
			"user_id":          user.(sharedModels.User).ID,
			"project_id":       participant.ProjectID,
		}).Warn("Unauthorized user attempting to delete participant")
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized!"})
		return
	}

	initializers.DB.Delete(&participant)
	logrus.WithFields(logrus.Fields{
		"user_id":    user.(sharedModels.User).ID,
		"project_id": participant.ProjectID,
	}).Info("Participation cancelled successfully")
	c.JSON(http.StatusOK, gin.H{"data": true})
}
