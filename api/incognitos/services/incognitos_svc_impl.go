package services

import (
	"fmt"
	"profile-api/api/incognitos/dto"
	"profile-api/api/incognitos/repositories"
	"profile-api/pkg/exceptions"
	"profile-api/pkg/mapper"
	
	emailDTO "profile-api/emails/dto"
	emails "profile-api/emails/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo     repositories.CompRepositories
	DB       *gorm.DB
	validate *validator.Validate
}

func NewComponentServices(compRepositories repositories.CompRepositories, db *gorm.DB, validate *validator.Validate) CompServices {
	return &CompServicesImpl{
		repo:     compRepositories,
		DB:       db,
		validate: validate,
	}
}

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.Incognitos) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	input := mapper.MapIncognitoInputToModel(data)
	input.UUID = uuid.NewString()

	err := s.repo.Create(ctx, s.DB, input)
	if err != nil {
		return err
	}

	incognitoData, err := s.repo.FindByUUID(ctx, s.DB, input.UUID)
	if err != nil {
		return err
	}

	go func() {
		err = emails.SendIncognitoEmail(emailDTO.EmailIncognites{
			Subject:       "Incognito Message",
			SentDate:      incognitoData.CreatedAt.Format("January 02, 2006 15:04 MST"),
			MessageID:     "MSG-" + incognitoData.UUID,
			MessageBody:   incognitoData.Message,
			SecurityLevel: "Confidential",
		})
		if err != nil {
			fmt.Println("Error sending email:", err)
		}
	}()

	return nil
}
