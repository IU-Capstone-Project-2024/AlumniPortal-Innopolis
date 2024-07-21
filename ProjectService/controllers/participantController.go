package controllers

import (
	"alumniportal.com/shared/helpers"
	"alumniportal.com/shared/initializers"
	sharedModels "alumniportal.com/shared/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ParticipantInput struct {
	TargetID  uint `json:"target" binding:"required"`
	IsProject bool `json:"is_project" binding:"required"`
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
	var participant sharedModels.Participant
	if input.IsProject {
		participant = sharedModels.Participant{
			UserID:    user.(sharedModels.User).ID,
			ProjectID: input.TargetID,
			Status:    helpers.Unverified,
			IsProject: true,
		}
	} else {
		participant = sharedModels.Participant{
			UserID:    user.(sharedModels.User).ID,
			EventID:   input.TargetID,
			Status:    helpers.Unverified,
			IsProject: false,
		}
	}

	if err := initializers.DB.Create(&participant).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to add participant")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if participant.IsProject {
		logrus.WithFields(logrus.Fields{
			"user_id":    user.(sharedModels.User).ID,
			"project_id": participant.ProjectID,
			"add_id":     participant.ID,
		}).Info("Participant added successfully")
		c.JSON(http.StatusOK, participant)
	} else {
		logrus.WithFields(logrus.Fields{
			"user_id":  user.(sharedModels.User).ID,
			"event_id": participant.ProjectID,
			"add_id":   participant.ID,
		}).Info("Participant added successfully")
		c.JSON(http.StatusOK, participant)
	}

}

func RemoveParticipant(c *gin.Context) {
	var participant sharedModels.Participant

	user, exists := c.Get("user")

	if !exists {
		logrus.Warn("User not authenticated for Removing Participant")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	if user.(sharedModels.User).Role == helpers.Admin {
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

	if participant.UserID != user.(sharedModels.User).ID && user.(sharedModels.User).Role != helpers.Admin {
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

func ApproveParticipant(c *gin.Context) {
	updateParticipantStatus(c, helpers.Accepted)
}

func DeclineParticipant(c *gin.Context) {
	updateParticipantStatus(c, helpers.Declined)
}

func updateParticipantStatus(c *gin.Context, status helpers.VerificationStatus) {
	var participant sharedModels.Participant
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&participant).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"participant_id": c.Param("id"),
			"error":          err.Error(),
		}).Error("Participant not found for updating status")
		c.JSON(http.StatusNotFound, gin.H{"error": "Participant not found"})
		return
	}

	participant.Status = status
	if err := initializers.DB.Save(&participant).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"participant_id": participant.ID,
			"error":          err.Error(),
		}).Error("Failed to update participant status")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"participant_id": participant.ID,
		"status":         status,
	}).Info("Participant status updated successfully")
	c.JSON(http.StatusOK, participant)
}
