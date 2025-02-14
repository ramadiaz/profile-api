package mapper

import (
	"profile-api/api/treeurls/dto"
	"profile-api/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapTreeURLInputToModel(input dto.TreeURLs) models.TreeURLs {
	var output models.TreeURLs

	mapstructure.Decode(input, &output)
	return output
}

func MapTreeURLModelToOutput(input models.TreeURLs) dto.TreeURLOutput {
	var output dto.TreeURLOutput

	mapstructure.Decode(input, &output)
	return output
}
