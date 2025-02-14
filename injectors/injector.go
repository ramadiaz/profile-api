// go:build wireinject
// go:build wireinject
//go:build wireinject
// +build wireinject

package injectors

import (
	incognitoControllers "profile-api/api/incognitos/controllers"
	incognitoRepositories "profile-api/api/incognitos/repositories"
	incognitoServices "profile-api/api/incognitos/services"

	likeControllers "profile-api/api/likes/controllers"
	likeRepositories "profile-api/api/likes/repositories"
	likeServices "profile-api/api/likes/services"
	
	treeControllers "profile-api/api/treeurls/controllers"
	treeRepositories "profile-api/api/treeurls/repositories"
	treeServices "profile-api/api/treeurls/services"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var incognitoFeatureSet = wire.NewSet(
	incognitoRepositories.NewComponentRepository,
	incognitoServices.NewComponentServices,
	incognitoControllers.NewCompController,
)

var likeFeatureSet = wire.NewSet(
	likeRepositories.NewComponentRepository,
	likeServices.NewComponentServices,
	likeControllers.NewCompController,
)

var treeFeatureSet = wire.NewSet(
	treeRepositories.NewComponentRepository,
	treeServices.NewComponentServices,
	treeControllers.NewCompController,
)

func InitializeIncognitoController(db *gorm.DB, validate *validator.Validate) incognitoControllers.CompControllers {
	wire.Build(incognitoFeatureSet)
	return nil
}

func InitializeLikeController(db *gorm.DB, validate *validator.Validate) likeControllers.CompControllers {
	wire.Build(likeFeatureSet)
	return nil
}

func InitializeTreeController(db *gorm.DB, validate *validator.Validate) treeControllers.CompControllers {
	wire.Build(treeFeatureSet)
	return nil
}
