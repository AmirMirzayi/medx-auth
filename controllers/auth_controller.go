package controllers

import (
	"medx/auth/domain/auth"
	"medx/auth/models"
	"medx/auth/services"
	"medx/auth/util"
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
	var request models.User
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}
	request.ID = primitive.NewObjectID()
	request.Password = util.Hash(request.Password)
	response, err := services.AuthService.Register(request)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.IndentedJSON(
		http.StatusOK,
		response,
	)
}

func (*authController) Login(ctx *gin.Context) {
	var (
		request *auth.LoginRequestBody
		user    *models.User
	)

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	user, err = services.AuthService.Login(request)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if user != nil {
		user.Password = ""
		ctx.IndentedJSON(
			http.StatusOK,
			gin.H{"token": "something hashed", "user": user},
		)
		return
	}

	ctx.String(http.StatusNotFound, "User or Password incorrect!")
}
