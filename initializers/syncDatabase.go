package initializers

import (
	"github.com/goJwt/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
