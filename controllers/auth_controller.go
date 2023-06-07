package controllers

import (
	"medx/auth/domain/auth"
	"medx/auth/models"
	"medx/auth/services"
	"medx/auth/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
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

	pwd, err := util.Hash(request.Password)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	request.ID = primitive.NewObjectID()
	request.Password = pwd
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

	// request.Password = pwd
	user, err = services.AuthService.Login(request)

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	if user != nil {
		err := util.Verify(user.Password, request.Password)

		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			ctx.String(http.StatusNotFound, "User or Password incorrect!")
			return
		}

		user.Password = ""
		id := user.ID
		token, err := util.GenerateToken(id.(primitive.ObjectID).String())
		if err != nil {
			ctx.String(http.StatusInternalServerError, "cant make jwt token")
			return
		}
		ctx.IndentedJSON(
			http.StatusOK,
			gin.H{
				"token": token,
				"user":  user,
			},
		)
		return
	}
}
