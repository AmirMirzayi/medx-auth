package auth

import "medx/auth/models"

type LoginRequestBody struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseBody struct {
	Token string `json:"token"`
	User models.User `json:"user"`
}
