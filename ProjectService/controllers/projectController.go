package controllers

import (
	"alumniportal.com/shared/helpers"
	"alumniportal.com/shared/initializers"
	sharedModels "alumniportal.com/shared/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ProjectInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func CreateProject(c *gin.Context) {
	var input ProjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to bind JSON for CreateProject")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := c.Get("user")
	projectRequest := sharedModels.Project{
		FounderID:   user.(sharedModels.User).ID,
		Name:        input.Name,
		Description: input.Description,
		Status:      helpers.Unverified,
	}

	if err := initializers.DB.Create(&projectRequest).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to create project")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id":    user.(sharedModels.User).ID,
		"project_id": projectRequest.ID,
	}).Info("Project created successfully")
	c.JSON(http.StatusOK, projectRequest)
}

func DeleteProject(c *gin.Context) {
	var project sharedModels.Project

	user, exists := c.Get("user")

	if !exists {
		logrus.Warn("User not authenticated for project removal")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	if user.(sharedModels.User).IsAdmin == true {
		if err := initializers.DB.Where("id = ?", c.Param("id")).First(&project).Error; err != nil {
			logrus.WithFields(logrus.Fields{
				"project_id": c.Param("id"),
				"error":      err.Error(),
			}).Error("Project not found for project removal")
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}
	} else {
		if err := initializers.DB.Where("id = ? AND status = ?", c.Param("id"), helpers.Unverified).First(&project).Error; err != nil {
			logrus.WithFields(logrus.Fields{
				"project_id": c.Param("id"),
				"error":      err.Error(),
			}).Error("Project not found or already accepted/declined for DeleteProject")
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found or already accepted/declined"})
			return
		}
	}

	if project.FounderID != user.(sharedModels.User).ID && user.(sharedModels.User).IsAdmin == false {
		logrus.WithFields(logrus.Fields{
			"user_id":    user.(sharedModels.User).ID,
			"project_id": project.ID,
		}).Warn("Unauthorized user attempting to delete project")
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized: User is not the owner of this project"})
		return
	}

	initializers.DB.Delete(&project)
	logrus.WithFields(logrus.Fields{
		"user_id":    user.(sharedModels.User).ID,
		"project_id": project.ID,
	}).Info("Project deleted successfully")
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func UpdateProject(c *gin.Context) {
	var project sharedModels.Project
	if err := initializers.DB.Where("id = ? AND status = ?", c.Param("id"), helpers.Unverified).First(&project).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"project_id": c.Param("id"),
			"error":      err.Error(),
		}).Error("Project not found or already accepted/declined for UpdateProject")
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found or already accepted/declined"})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		logrus.Warn("User not authenticated for UpdateProject")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	if project.FounderID != user.(sharedModels.User).ID {
		logrus.WithFields(logrus.Fields{
			"user_id":    user.(sharedModels.User).ID,
			"project_id": project.ID,
		}).Warn("Unauthorized user attempting to update project")
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized: User is not the owner of this project"})
		return
	}

	var input ProjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to bind JSON for UpdateProject")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateData := sharedModels.Project{
		Name:        input.Name,
		Description: input.Description,
	}

	if err := initializers.DB.Model(&project).Updates(updateData).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to update project")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id":    user.(sharedModels.User).ID,
		"project_id": project.ID,
	}).Info("Project updated successfully")
	c.JSON(http.StatusOK, project)
}

func UpdateProjectAdmin(c *gin.Context) {
	var project sharedModels.Project
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&project).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"project_id": c.Param("id"),
			"error":      err.Error(),
		}).Error("Project not found for updating")
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		logrus.Warn("User not authenticated for project update")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var input ProjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to bind JSON for project update")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateData := sharedModels.Project{
		Name:        input.Name,
		Description: input.Description,
	}

	if err := initializers.DB.Model(&project).Updates(updateData).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to update project")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id":    user.(sharedModels.User).ID,
		"project_id": project.ID,
	}).Info("Project updated successfully by admin")
	c.JSON(http.StatusOK, project)
}

func ApproveProject(c *gin.Context) {
	updateProjectVerificationStatus(c, helpers.Accepted)
}

func DeclineProject(c *gin.Context) {
	updateProjectVerificationStatus(c, helpers.Declined)
}

func updateProjectVerificationStatus(c *gin.Context, status helpers.VerificationStatus) {
	var project sharedModels.Project
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&project).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"project_id": c.Param("id"),
			"error":      err.Error(),
		}).Error("Project not found for update Project Status")
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	project.Status = status
	if err := initializers.DB.Save(&project).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"project_id": project.ID,
			"error":      err.Error(),
		}).Error("Failed to update project verification status")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"project_id": project.ID,
		"status":     status,
	}).Info("Project verification status updated successfully")
	c.JSON(http.StatusOK, project)
}

func GetCurrentUserProjects(c *gin.Context) {
	user, _ := c.Get("user")
	var projects []sharedModels.Project
	if err := initializers.DB.Where("founder_id = ?", user.(sharedModels.User).ID).Preload("User").Find(&projects).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"user_id": user.(sharedModels.User).ID,
			"error":   err.Error(),
		}).Error("Failed to get current user projects")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id": user.(sharedModels.User).ID,
		"count":   len(projects),
	}).Info("Fetched current user projects successfully")
	c.JSON(http.StatusOK, projects)
}

func GetUnverifiedProjects(c *gin.Context) {
	var projects []sharedModels.Project
  
	if err := initializers.DB.Where("status = ?", helpers.Unverified).Preload("User").Find(&projects).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to get unverified projects")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"count": len(projects),
	}).Info("Fetched unverified projects successfully")
	c.JSON(http.StatusOK, projects)
}

func GetProjectAdmin(c *gin.Context) {
	var project sharedModels.Project

	if err := initializers.DB.Where("status = ? AND id = ?", helpers.Unverified, c.Param("id")).Preload("User").First(&project).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"project_id": c.Param("id"),
			"error":      err.Error(),
		}).Error("Failed to get project by admin")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"project_id": project.ID,
	}).Info("Fetched project by admin successfully")
	c.JSON(http.StatusOK, project)
}

func GetProject(c *gin.Context) {
	var project sharedModels.Project

	user, _ := c.Get("user")

	if err := initializers.DB.Where("id = ? AND founder_id = ?", c.Param("id"), user.(sharedModels.User).ID).Preload("User").First(&project).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"user_id":    user.(sharedModels.User).ID,
			"project_id": c.Param("id"),
			"error":      err.Error(),
		}).Error("Failed to get project")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id":    user.(sharedModels.User).ID,
		"project_id": project.ID,
	}).Info("Fetched project successfully")
	c.JSON(http.StatusOK, project)
}
