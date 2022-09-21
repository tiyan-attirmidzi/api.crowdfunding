package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (c *authController) CheckEmailAvailability(ctx *gin.Context) {

	var payload dto.CheckEmail

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		res := helpers.ResponseJSON(
			"Email Checking Failed!",
			http.StatusUnprocessableEntity,
			"error",
			gin.H{"errors": helpers.InputValidation(err)},
		)
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	isEmailAvailable, err := c.authService.IsEmailAvailable(payload)

	if err != nil {
		res := helpers.ResponseJSON(
			"Email Address Has Been Registered!",
			http.StatusUnprocessableEntity,
			"error",
			gin.H{"is_available": isEmailAvailable},
		)
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	res := helpers.ResponseJSON("Email Address Can Be Used!", http.StatusOK, "success", gin.H{"is_available": isEmailAvailable})
	ctx.JSON(http.StatusOK, res)

}

func (c *authController) UploadAvatar(ctx *gin.Context) {

	file, err := ctx.FormFile("avatar")

	if err != nil {
		res := helpers.ResponseJSON(
			"Failed to Upload Avatar image!",
			http.StatusBadRequest,
			"error",
			gin.H{"is_uploaded": false},
		)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	fmt.Println("INI FILE =>", file.Header)

	// retrieve file information
	extension := filepath.Ext(file.Filename)
	// path file and generate filename
	pathFile := "uploads/images/" + uuid.New().String() + extension

	if err = ctx.SaveUploadedFile(file, pathFile); err != nil {
		res := helpers.ResponseJSON(
			"Failed to Upload Avatar image!",
			http.StatusBadRequest,
			"error",
			err.Error(),
		)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	// TODO: change to dynamic latter
	userID := 1

	_, err = c.authService.SaveAvatar(userID, pathFile)

	if err != nil {
		res := helpers.ResponseJSON(
			"Failed to Upload Avatar image!",
			http.StatusBadRequest,
			"error",
			gin.H{"is_uploaded": false},
		)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := helpers.ResponseJSON(
		"Uploaded Avatar Successfully!",
		http.StatusOK,
		"error",
		gin.H{"is_uploaded": true},
	)
	ctx.JSON(http.StatusOK, res)

}
