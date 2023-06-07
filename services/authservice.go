package services

import (
	"medx/auth/domain/auth"
	"medx/auth/models"
)

var AuthService authServiceInterface = &authService{}

type authServiceInterface interface {
	Register(models.User) (*auth.ResponseBody, error)
	Login(*auth.LoginRequestBody) (*models.User, error)
}

type authService struct{}

func (*authService) Register(user models.User) (*auth.ResponseBody, error) {

	db, err := InitDB()

	if err != nil {
		return nil, err
	}

	res, err := db.Create(user)

	if err != nil {
		return nil, err
	}

	return &auth.ResponseBody{
		Token: "something hashed!",
		User:  res,
	}, nil
}

func (*authService) Login(credential *auth.LoginRequestBody) (*models.User, error) {
	db, err := InitDB()
	if err != nil {
		return nil, err
	}

	res, err := db.Find(*credential)

	if err != nil {
		return nil, err
	}

	return res, nil
}
