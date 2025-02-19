package main

import (
	"profile-api/models"
	"profile-api/pkg/config"
)

func main() {
	db := config.InitDB()

	if err := db.Exec(`DO $$ 
    		BEGIN
    		    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'featured_type') THEN
    		        CREATE TYPE featured_type AS ENUM ('hot', 'featured');
    		    END IF;
    		END $$;`).
		Error; err != nil {
		panic("failed to create enum type: " + err.Error())
	}

	err := db.AutoMigrate(
		&models.Client{},
		&models.Incognitos{},
		&models.Likes{},
		&models.TreeURLs{},
		&models.Blogs{},
		&models.BlogTags{},
		&models.FeaturedBlogs{},
		&models.Files{},
	)
	if err != nil {
		panic("failed to migrate models: " + err.Error())
	}
}
