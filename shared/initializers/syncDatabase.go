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
	if DB.AutoMigrate(&models.Participant{}) != nil {
		panic("Database models.Participant migration failed")
	}
	if DB.AutoMigrate(&models.Event{}) != nil {
		panic("Database models.Event migration failed")
	}
	if DB.AutoMigrate(&models.Donation{}) != nil {
		panic("Database models.Donation migration failed")
	}
	if DB.AutoMigrate(&models.Volunteer{}) != nil {
		panic("Database models.Volunteer migration failed")
	}

	logrus.Info("Database migration completed!")
}
