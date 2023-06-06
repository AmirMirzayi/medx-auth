package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var AuthController authControllerInterface = &authController{}

type authControllerInterface interface {
	Register(*gin.Context)
	Login(*gin.Context)
}
type authController struct{}

func (*authController) Register(ctx *gin.Context) {
	ctx.String(http.StatusOK, "wait, who are you?")
}

func (*authController) Login(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hi amir, welcome :)")
}
