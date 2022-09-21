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
	userService services.UserService
}

func NewAuthController(authService services.AuthService, userService services.UserService) *authController {
	return &authController{authService, userService}
}

func (c *authController) SignUp(ctx *gin.Context) {

	var payload dto.SignUp

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		res := helpers.ResponseJSON(
			"user sign up failed!",
			http.StatusUnprocessableEntity,
			"error",
			gin.H{"errors": helpers.InputValidation(err)},
		)
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	// check email availability
	_, err := c.userService.IsEmailAvailable(payload.Email)

	if err != nil {
		res := helpers.ResponseJSON(
			"email address has been registered!",
			http.StatusUnprocessableEntity,
			"error",
			nil,
		)
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	signedUp, err := c.authService.SignUp(payload)

	if err != nil {
		res := helpers.ResponseJSON("user signed up failed!", http.StatusBadRequest, "error", gin.H{"errors": err.Error()})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	token, err := c.authService.GenerateToken(signedUp)

	if err != nil {
		res := helpers.ResponseJSON("user signed up failed!", http.StatusBadRequest, "error", gin.H{"errors": err.Error()})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := helpers.ResponseJSON("user signed up successfully!", http.StatusCreated, "success", dto.FormatAuth(signedUp, token))
	ctx.JSON(http.StatusCreated, res)

}

func (c *authController) SignIn(ctx *gin.Context) {

	var payload dto.SignIn

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		res := helpers.ResponseJSON(
			"user sign in failed!",
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
			"user sign in failed!",
			http.StatusBadRequest,
			"error",
			gin.H{"errors": err.Error()},
		)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	token, err := c.authService.GenerateToken(signedIn)

	if err != nil {
		res := helpers.ResponseJSON("user sign up failed!", http.StatusBadRequest, "error", gin.H{"errors": err.Error()})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := helpers.ResponseJSON("user signed in successfully!", http.StatusOK, "success", dto.FormatAuth(signedIn, token))
	ctx.JSON(http.StatusOK, res)

}
