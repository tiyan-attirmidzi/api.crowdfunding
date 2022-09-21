package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiyan-attirmidzi/api.crowdfunding/entities/dto"
	"github.com/tiyan-attirmidzi/api.crowdfunding/helpers"
	"github.com/tiyan-attirmidzi/api.crowdfunding/services"
)

type authController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *authController {
	return &authController{authService}
}

func (c *authController) SignUp(ctx *gin.Context) {

	var payload dto.SignUp

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		res := helpers.ResponseJSON(
			"User Sign Up Failed!",
			http.StatusUnprocessableEntity,
			"error",
			gin.H{"errors": helpers.InputValidation(err)},
		)
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	newUser, err := c.authService.SignUp(payload)

	if err != nil {
		res := helpers.ResponseJSON("User Signed Up Failed!", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	formatter := dto.FormatUser(newUser, "TokenHashToken")

	res := helpers.ResponseJSON("User Signed Up Successfully!", http.StatusCreated, "success", formatter)
	ctx.JSON(http.StatusCreated, res)

}

func (c *authController) SignIn(ctx *gin.Context) {

	var payload dto.SignIn

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		res := helpers.ResponseJSON(
			"User Sign In Failed!",
			http.StatusUnprocessableEntity,
			"error",
			gin.H{"errors": helpers.InputValidation(err)},
		)
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	signedIn, err := c.authService.SignIn(payload)

	if err != nil {
		res := helpers.ResponseJSON(
			"User Sign In Failed!",
			http.StatusUnprocessableEntity,
			"error",
			gin.H{"errors": err.Error()},
		)
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	formatter := dto.FormatUser(signedIn, "TokenHashToken")

	res := helpers.ResponseJSON("User Signed In Successfully!", http.StatusOK, "success", formatter)
	ctx.JSON(http.StatusOK, res)

}
