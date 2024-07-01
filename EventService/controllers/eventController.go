package controllers

import (
	"alumniportal.com/shared/helpers"
	"alumniportal.com/shared/initializers"
	sharedModels "alumniportal.com/shared/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/datatypes"
	"net/http"
)

type EventInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Datetime    string `json:"date" binding:"required"`
	Duration    int    `json:"duration" binding:"required"`
}

func CreateEvent(c *gin.Context) {
	var input EventInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to bind JSON for event creation")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := c.Get("user")
	eventRequest := sharedModels.Event{
		FounderID:   user.(sharedModels.User).ID,
		Name:        input.Name,
		Description: input.Description,
		Date:        datatypes.Date(helpers.Convert(input.Datetime)),
		Duration:    input.Duration,
		Status:      helpers.Unverified,
	}

	if err := initializers.DB.Create(&eventRequest).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to create event")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id":  user.(sharedModels.User).ID,
		"event_id": eventRequest.ID,
	}).Info("Event created successfully")
	c.JSON(http.StatusOK, eventRequest)
}

func DeleteEvent(c *gin.Context) {
	var event sharedModels.Event

	user, exists := c.Get("user")

	if !exists {
		logrus.Warn("User not authenticated for event removal")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	if user.(sharedModels.User).IsAdmin == true {
		if err := initializers.DB.Where("id = ?", c.Param("id")).First(&event).Error; err != nil {
			logrus.WithFields(logrus.Fields{
				"event_id": c.Param("id"),
				"error":    err.Error(),
			}).Error("Event not found for event removal")
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			return
		}
	} else {
		if err := initializers.DB.Where("id = ? AND status = ?", c.Param("id"), helpers.Unverified).First(&event).Error; err != nil {
			logrus.WithFields(logrus.Fields{
				"event_id": c.Param("id"),
				"error":    err.Error(),
			}).Error("Event not found or already accepted/declined")
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found or already accepted/declined"})
			return
		}
	}

	if event.FounderID != user.(sharedModels.User).ID && user.(sharedModels.User).IsAdmin == false {
		logrus.WithFields(logrus.Fields{
			"user_id":  user.(sharedModels.User).ID,
			"event_id": event.ID,
		}).Warn("Unauthorized user attempting to delete event")
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized: User is not the owner of this event"})
		return
	}

	initializers.DB.Delete(&event)
	logrus.WithFields(logrus.Fields{
		"user_id":  user.(sharedModels.User).ID,
		"event_id": event.ID,
	}).Info("Event deleted successfully")
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func UpdateEvent(c *gin.Context) {
	var event sharedModels.Event

	if err := initializers.DB.Where("id = ? AND status = ?", c.Param("id"), helpers.Unverified).First(&event).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event_id": c.Param("id"),
			"error":    err.Error(),
		}).Error("Event not found or already accepted/declined")
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found or already accepted/declined"})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		logrus.Warn("User not authenticated for updating event")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	if event.FounderID != user.(sharedModels.User).ID {
		logrus.WithFields(logrus.Fields{
			"user_id":  user.(sharedModels.User).ID,
			"event_id": event.ID,
		}).Warn("Unauthorized user attempting to update event")
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized: User is not the owner of this event"})
		return
	}

	var input EventInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to bind JSON for event")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateData := sharedModels.Event{
		Name:        input.Name,
		Description: input.Description,
		Date:        datatypes.Date(helpers.Convert(input.Datetime)),
		Duration:    input.Duration,
	}

	if err := initializers.DB.Model(&event).Updates(updateData).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to update event")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id":  user.(sharedModels.User).ID,
		"event_id": event.ID,
	}).Info("Event updated successfully")
	c.JSON(http.StatusOK, event)
}

func UpdateEventAdmin(c *gin.Context) {
	var event sharedModels.Event

	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&event).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event_id": c.Param("id"),
			"error":    err.Error(),
		}).Error("Event not found for updating")
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		logrus.Warn("User not authenticated for event update")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var input EventInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to bind JSON for event update")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateData := sharedModels.Event{
		Name:        input.Name,
		Description: input.Description,
		Date:        datatypes.Date(helpers.Convert(input.Datetime)),
		Duration:    input.Duration,
	}

	if err := initializers.DB.Model(&event).Updates(updateData).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to update event")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id":  user.(sharedModels.User).ID,
		"event_id": event.ID,
	}).Info("Event updated successfully by admin")
	c.JSON(http.StatusOK, event)
}

func ApproveEvent(c *gin.Context) {
	updateEventVerificationStatus(c, helpers.Accepted)
}

func DeclineEvent(c *gin.Context) {
	updateEventVerificationStatus(c, helpers.Declined)
}

func updateEventVerificationStatus(c *gin.Context, status helpers.VerificationStatus) {
	var event sharedModels.Event
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&event).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event_id": c.Param("id"),
			"error":    err.Error(),
		}).Error("Event not found for update Event Status")
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	event.Status = status
	if err := initializers.DB.Save(&event).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event_id": event.ID,
			"error":    err.Error(),
		}).Error("Failed to update event verification status")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"event_id": event.ID,
		"status":   status,
	}).Info("Event verification status updated successfully")
	c.JSON(http.StatusOK, event)
}

func GetCurrentUserEvents(c *gin.Context) {
	user, _ := c.Get("user")

	var events []sharedModels.Event

	if err := initializers.DB.Where("founder_id = ?", user.(sharedModels.User).ID).Preload("User").Find(&events).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"user_id": user.(sharedModels.User).ID,
			"error":   err.Error(),
		}).Error("Failed to get current user events")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id": user.(sharedModels.User).ID,
		"count":   len(events),
	}).Info("Fetched current user events successfully")
	c.JSON(http.StatusOK, events)
}

func GetUnverifiedEvents(c *gin.Context) {
	var events []sharedModels.Event

	if err := initializers.DB.Where("status = ?", helpers.Unverified).Preload("User").Find(&events).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to get unverified events")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"count": len(events),
	}).Info("Fetched unverified events successfully")
	c.JSON(http.StatusOK, events)
}

func GetEventAdmin(c *gin.Context) {
	var event sharedModels.Event

	if err := initializers.DB.Where("status = ? AND id = ?", helpers.Unverified, c.Param("id")).Preload("User").First(&event).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"event_id": c.Param("id"),
			"error":    err.Error(),
		}).Error("Failed to get event by admin")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"event_id": event.ID,
	}).Info("Fetched event by admin successfully")
	c.JSON(http.StatusOK, event)
}

func GetEvent(c *gin.Context) {
	var event sharedModels.Event

	user, _ := c.Get("user")

	if err := initializers.DB.Where("id = ? AND founder_id = ?", c.Param("id"), user.(sharedModels.User).ID).Preload("User").First(&event).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"user_id":  user.(sharedModels.User).ID,
			"event_id": c.Param("id"),
			"error":    err.Error(),
		}).Error("Failed to get event")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id":  user.(sharedModels.User).ID,
		"event_id": event.ID,
	}).Info("Fetched event successfully")
	c.JSON(http.StatusOK, event)
}
