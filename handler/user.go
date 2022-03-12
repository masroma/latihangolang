package handler

import (
	"latihangolang/helper"
	"latihangolang/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisteruserInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Register failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Register failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(newUser, "tokentokentoken")
	response := helper.APIResponse("Register successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// token, err := h.authService.GenerateToken(loggedinUser.ID)
	// if err != nil {
	// 	errorMessage := gin.H{"errors" : err.Error()}
	// 	response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
	// 	c.JSON(http.StatusUnprocessableEntity, response)
	// 	return
	// }

	token := "tokentokentoken"

	formatter := user.FormatUser(loggedinUser, token)

	response := helper.APIResponse("Account login", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) CheckEmailAvailable(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Check Email Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	isEmailAvailable, err := h.userService.EmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server Error"}
		response := helper.APIResponse("Check Email Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	var metaMessage string
	if isEmailAvailable {
		metaMessage = "Email available"
	} else {
		metaMessage = "Email has been register"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) UploadPhotoProfile(c *gin.Context) {
	file, err := c.FormFile("photo")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("failed to upload photo profile", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := "images/" + file.Filename
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("failed to upload photo profile", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID := 5
	_, err = h.userService.UpdateAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("failed to upload photo profile", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("success to upload photo profile", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)

}
