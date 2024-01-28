package initializers

import "github.com/reynaldineo/Go-JWT-Auth/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}