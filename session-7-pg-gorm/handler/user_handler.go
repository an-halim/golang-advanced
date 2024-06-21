package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/an-halim/golang-advanced/session-7-pg-gorm/entity"
	"github.com/an-halim/golang-advanced/session-7-pg-gorm/response"
	"github.com/an-halim/golang-advanced/session-7-pg-gorm/service"
	"github.com/gin-gonic/gin"
)

// IUserHandler mendefinisikan interface untuk handler user
type IUserHandler interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetAllUsers(c *gin.Context)
}

type UserHandler struct {
	userService service.IUserService
}

// NewUserHandler membuat instance baru dari UserHandler
func NewUserHandler(userService service.IUserService) IUserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// CreateUser menghandle permintaan untuk membuat user baru
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		errMsg := err.Error()
		errMsg = convertUserMandatoryFieldErrorString(errMsg)
		apiResponse := response.Failed{
			Status:  "failed",
			Message: errMsg,
		}
		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	createdUser, err := h.userService.CreateUser(c.Request.Context(), &user)
	if err != nil {
		apiResponse := response.Failed{
			Status:  "failed",
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	apiResponse := response.Success{
		Status:  "success",
		Message: "User created successfully",
		Data:    createdUser,
	}

	c.JSON(http.StatusCreated, apiResponse)
}

// GetUser menghandle permintaan untuk mendapatkan user berdasarkan ID
func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		apiResponse := response.Failed{
			Status:  "failed",
			Message: "Invalid ID",
		}
		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	user, err := h.userService.GetUserByID(c.Request.Context(), id)
	if (err != nil) || (user.Name == "") {
		apiResponse := response.Failed{
			Status:  "failed",
			Message: "User not found",
		}
		c.JSON(http.StatusNotFound, apiResponse)
		return
	}

	apiResponse := response.Success{
		Status:  "success",
		Message: "User found",
		Data:    user,
	}

	c.JSON(http.StatusOK, apiResponse)
}

// UpdateUser menghandle permintaan untuk mengupdate informasi user
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {

		apiResponse := response.Failed{
			Status:  "failed",
			Message: "Invalid ID",
		}
		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		errMsg := err.Error()
		errMsg = convertUserMandatoryFieldErrorString(errMsg)
		apiResponse := response.Failed{
			Status:  "failed",
			Message: errMsg,
		}
		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	updatedUser, err := h.userService.UpdateUser(c.Request.Context(), id, user)
	if err != nil {
		apiResponse := response.Failed{
			Status:  "failed",
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, apiResponse)
		return
	}

	apiResponse := response.Success{
		Status:  "success",
		Message: "User updated successfully",
		Data:    updatedUser,
	}

	c.JSON(http.StatusOK, apiResponse)
}

// DeleteUser menghandle permintaan untuk menghapus user
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		apiResponse := response.Failed{
			Status:  "failed",
			Message: "Invalid ID",
		}
		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	if err := h.userService.DeleteUser(c.Request.Context(), id); err != nil {
		apiResponse := response.Failed{
			Status:  "failed",
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, apiResponse)
		return
	}

	apiResponse := response.Success{
		Status:  "success",
		Message: "User deleted successfully",
	}

	c.JSON(http.StatusOK, apiResponse)
}

// GetAllUsers menghandle permintaan untuk mendapatkan semua user
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	pageSize := c.DefaultQuery("size", "10")
	currentPage := c.DefaultQuery("page", "1")

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		apiResponse := response.Failed{
			Status:  "failed",
			Message: "Invalid page size",
		}
		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	currentPageInt, err := strconv.Atoi(currentPage)
	if err != nil {
		apiResponse := response.Failed{
			Status:  "failed",
			Message: "Invalid page number",
		}
		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	offset := pageSizeInt * (currentPageInt - 1)

	users, err := h.userService.GetAllUsers(c.Request.Context(), pageSizeInt, offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	apiResponse := response.GetAll{
		Status:      "success",
		Data:        users,
		PageSize:    pageSizeInt,
		CurrentPage: currentPageInt,
	}

	c.JSON(http.StatusOK, apiResponse)
}

func convertUserMandatoryFieldErrorString(oldErrorMsg string) string {
	switch {
	case strings.Contains(oldErrorMsg, "'Name' failed on the 'required' tag"):
		return "name is mandatory"
	case strings.Contains(oldErrorMsg, "'Email' failed on the 'required' tag"):
		return "email is mandatory"
	}
	return oldErrorMsg
}
