package initializers

import (
	"alumniportal.com/shared/models"
	"github.com/sirupsen/logrus"
)

func SyncDatabase() {
	if DB.AutoMigrate(&models.User{}) != nil {
		panic("Database models.User migration failed")
	}
	logrus.Info("Database models.User migration completed!")
}
