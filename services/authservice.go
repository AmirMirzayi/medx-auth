package services

import (
	"medx/auth/domain/auth"
	"medx/auth/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var AuthService authServiceInterface = &authService{}

type authServiceInterface interface {
	Register(*models.User) (auth.LoginResponseBody, error)
	Login(*auth.LoginRequestBody) (bool, error)
}
type authService struct{}

func (*authService) Register(user *models.User) (auth.LoginResponseBody, error) {
	return auth.LoginResponseBody{
		Token: "something hashed!",
		User: models.User{
			ID:          primitive.NewObjectID(),
			UserName:    "Amir",
			Password:    "Mirzaei",
			FirstName:   "Amir",
			Address:     "Mashhad",
			PhoneNumber: "09105599950",
		},
	}, nil
}

func (*authService) Login(credential *auth.LoginRequestBody) (bool, error) {
	if credential.UserName == "amir" {
		return true, nil
	}
	return false, nil
}
