package mapper

import (
	"profile-api/api/incognitos/dto"
	"profile-api/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapIncognitoInputToModel(input dto.Incognitos) models.Incognitos {
	var incognito models.Incognitos

	mapstructure.Decode(input, &incognito)
	return incognito
}
