package main

import (
	"log"
	"os"
	"profile-api/pkg/config"
	"profile-api/pkg/middleware"
	"profile-api/routers"
	"time"

	internalRouters "profile-api/internal/routers"

	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth/v7/limiter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	config.InitEnvCheck()

	port := os.Getenv("PORT")
	environment := os.Getenv("ENVIRONMENT")

	r := gin.New()
	r.Use(gin.Logger())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))

	db := config.InitDB()
	storage := config.InitStorage()
	validate := validator.New(validator.WithRequiredStructEnabled())
	lmt := tollbooth.NewLimiter(5, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Second})

	r.Use(middleware.ClientTracker(db))
	r.Use(middleware.GzipResponseMiddleware())
	r.Use(middleware.RateLimitMiddleware(lmt))

	internal := r.Group("/internal")
	internalRouters.InternalRouters(internal, db, storage, validate)

	api := r.Group("/api")
	routers.CompRouters(api, db, validate)

	var host string
	switch environment {
	case "development":
		host = "localhost"
	case "production":
		host = "0.0.0.0"
	default:
		panic("ENV ERROR: {ENVIRONMENT} UNKNOWN")
	}

	server := host + ":" + port
	err := r.Run(server)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}

	log.Println("Server started on port :" + port)
}
