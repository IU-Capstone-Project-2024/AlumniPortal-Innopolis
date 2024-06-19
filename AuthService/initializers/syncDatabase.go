package initializers

import (
	"AuthService/models"
)

func SyncDatabase() {
	if DB.AutoMigrate(&models.User{}) != nil {
		panic("Database models.User migration failed")
	}
}
