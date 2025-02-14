package main

import (
	"profile-api/pkg/config"
	"profile-api/models"
)

func main() {
	db := config.InitDB()

	err := db.AutoMigrate(
		&models.Client{},
		&models.Example{},
	)
	if err != nil {
		panic("failed to migrate models: " + err.Error())
	}
}
