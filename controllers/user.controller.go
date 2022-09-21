package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tiyan-attirmidzi/api.crowdfunding/entities/dto"
	"github.com/tiyan-attirmidzi/api.crowdfunding/helpers"
	"github.com/tiyan-attirmidzi/api.crowdfunding/services"
)

type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *userController {
	return &userController{userService}
}

func (c *userController) CheckEmailAvailability(ctx *gin.Context) {

	var payload dto.CheckExistingEmail

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

	isEmailAvailable, err := c.userService.IsEmailAvailable(payload.Email)

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

func (c *userController) UploadAvatar(ctx *gin.Context) {

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

	_, err = c.userService.SaveAvatar(userID, pathFile)

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
