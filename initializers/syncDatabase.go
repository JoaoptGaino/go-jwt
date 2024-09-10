package initializers

import "github.com/joaoptgaino/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
