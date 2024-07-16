package initializers

import (
	"alumniportal.com/shared/models"
	"fmt"
	"github.com/sirupsen/logrus"
)

func SyncDatabase() {
	migrateModel(&models.User{})
	migrateModel(&models.PassRequest{})
	migrateModel(&models.Project{})
	migrateModel(&models.Participant{})
	migrateModel(&models.Event{})
	migrateModel(&models.Donation{})
	migrateModel(&models.Volunteer{})

	logrus.Info("Database migration completed!")
}

func migrateModel(model interface{}) {
	if !DB.Migrator().HasTable(model) {
		if err := DB.AutoMigrate(model); err != nil {
			panic(fmt.Sprintf("Database migration for %T failed: %v", model, err))
		}
	} else {
		logrus.Infof("Skipping migration for %T - table already exists", model)
	}
}
