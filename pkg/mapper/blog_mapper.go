package mapper

import (
	"profile-api/api/blogs/dto"
	"profile-api/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapBlogInputToModel(input dto.Blogs) models.Blogs {
	var output models.Blogs

	mapstructure.Decode(input, &output)
	return output
}

func MapBlogModelToOutput(input models.Blogs) dto.BlogOutput {
	var output dto.BlogOutput

	mapstructure.Decode(input, &output)
	output.CreatedAt = input.CreatedAt
	output.UpdatedAt = input.UpdatedAt
	return output
}
