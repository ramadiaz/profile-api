package services

import (
	"net/http"
	"os"
	"profile-api/internal/auth/dto"
	"profile-api/pkg/exceptions"
	"profile-api/pkg/variables"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	DB       *gorm.DB
	validate *validator.Validate
}

func NewComponentServices(validate *validator.Validate) CompServices {
	return &CompServicesImpl{
		validate: validate,
	}
}

func (s *CompServicesImpl) Login(ctx *gin.Context, data dto.Login) (*string, *exceptions.Exception) {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return nil, exceptions.NewValidationException(validateErr)
	}

	ADMIN_USERNAME := os.Getenv("ADMIN_USERNAME")
	ADMIN_PASSWORD := os.Getenv("ADMIN_PASSWORD")

	if data.Username != ADMIN_USERNAME || data.Password != ADMIN_PASSWORD {
		return nil, exceptions.NewException(http.StatusUnauthorized, exceptions.ErrInvalidCredentials)
	}

	JWT_SECRET := os.Getenv("JWT_SECRET")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = ADMIN_USERNAME
	claims["role"] = variables.ADMIN

	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	secretKey := []byte(JWT_SECRET)
	tokenString, signErr := token.SignedString(secretKey)
	if signErr != nil {
		return nil, exceptions.NewException(http.StatusInternalServerError, exceptions.ErrTokenGenerate)
	}

	return &tokenString, nil
}
