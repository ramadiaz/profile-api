package main

import (
	"profile-api/models"
	"profile-api/pkg/config"
)

func main() {
	db := config.InitDB()

	err := db.AutoMigrate(
		&models.Client{},
		&models.Incognitos{},
		&models.Likes{},
	)
	if err != nil {
		panic("failed to migrate models: " + err.Error())
	}
}
