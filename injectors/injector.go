// go:build wireinject
// go:build wireinject
//go:build wireinject
// +build wireinject

package injectors

import (
	incognitoControllers "profile-api/api/incognitos/controllers"
	incognitoRepositories "profile-api/api/incognitos/repositories"
	incognitoServices "profile-api/api/incognitos/services"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var incognitoFeatureSet = wire.NewSet(
	incognitoRepositories.NewComponentRepository,
	incognitoServices.NewComponentServices,
	incognitoControllers.NewCompController,
)

func InitializeIncognitoController(db *gorm.DB, validate *validator.Validate) incognitoControllers.CompControllers {
	wire.Build(incognitoFeatureSet)
	return nil
}
