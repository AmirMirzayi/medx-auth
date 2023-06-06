package services

import (
	"medx/auth/domain/auth"
	"medx/auth/models"
)

var AuthService authServiceInterface = &authService{}

type authServiceInterface interface {
	Register(*models.User)
	Login(*auth.LoginRequestBody)
}
type authService struct{}

func (*authService) Register(user *models.User) {

}

func (*authService) Login(*auth.LoginRequestBody) {

}
