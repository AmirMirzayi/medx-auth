package controllers

import (
	"medx/auth/domain/auth"
	"medx/auth/models"
	"medx/auth/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var AuthController authControllerInterface = &authController{}

type authControllerInterface interface {
	Register(*gin.Context)
	Login(*gin.Context)
}

type authController struct{}

func (*authController) Register(ctx *gin.Context) {
	var request *models.User
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	request.ID = primitive.NewObjectID()
	response, err := services.AuthService.Register(request)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	response.User.Password = ""
	ctx.IndentedJSON(
		http.StatusOK,
		response,
	)
}

func (*authController) Login(ctx *gin.Context) {
	var request *auth.LoginRequestBody
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	isValid, err := services.AuthService.Login(request)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if isValid {
		ctx.String(http.StatusOK, "hi amir, welcome :)")
		return
	}

	ctx.String(http.StatusNotFound, "User or Password incorrect!")
}
