package initializers

import (
	"alumniportal.com/shared/models"
	"github.com/sirupsen/logrus"
)

func SyncDatabase() {
	if DB.AutoMigrate(&models.User{}) != nil {
		panic("Database models.User migration failed")
	}
	if DB.AutoMigrate(&models.PassRequest{}) != nil {
		panic("Database models.PassRequest migration failed")
	}
	if DB.AutoMigrate(&models.Project{}) != nil {
		panic("Database models.Project migration failed")
	}

	logrus.Info("Database models.User migration completed!")
}
