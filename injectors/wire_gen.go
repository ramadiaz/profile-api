// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injectors

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
	"profile-api/api/incognitos/controllers"
	"profile-api/api/incognitos/repositories"
	"profile-api/api/incognitos/services"
	controllers2 "profile-api/api/likes/controllers"
	repositories2 "profile-api/api/likes/repositories"
	services2 "profile-api/api/likes/services"
	controllers3 "profile-api/api/treeurls/controllers"
	repositories3 "profile-api/api/treeurls/repositories"
	services3 "profile-api/api/treeurls/services"
	"profile-api/storages"
)

// Injectors from injector.go:

func InitializeIncognitoController(db *gorm.DB, validate *validator.Validate) controllers.CompControllers {
	compRepositories := repositories.NewComponentRepository()
	compServices := services.NewComponentServices(compRepositories, db, validate)
	compControllers := controllers.NewCompController(compServices)
	return compControllers
}

func InitializeLikeController(db *gorm.DB, validate *validator.Validate) controllers2.CompControllers {
	compRepositories := repositories2.NewComponentRepository()
	compServices := services2.NewComponentServices(compRepositories, db, validate)
	compControllers := controllers2.NewCompController(compServices)
	return compControllers
}

func InitializeTreeController(db *gorm.DB, validate *validator.Validate) controllers3.CompControllers {
	compRepositories := repositories3.NewComponentRepository()
	memory := storages.NewMemory()
	compServices := services3.NewComponentServices(compRepositories, db, validate, memory)
	compControllers := controllers3.NewCompController(compServices)
	return compControllers
}

// injector.go:

var incognitoFeatureSet = wire.NewSet(repositories.NewComponentRepository, services.NewComponentServices, controllers.NewCompController)

var likeFeatureSet = wire.NewSet(repositories2.NewComponentRepository, services2.NewComponentServices, controllers2.NewCompController)

var treeFeatureSet = wire.NewSet(storages.NewMemory, repositories3.NewComponentRepository, services3.NewComponentServices, controllers3.NewCompController)
