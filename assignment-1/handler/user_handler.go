package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/an-halim/golang-advanced/assignment-1/request"
	"github.com/an-halim/golang-advanced/assignment-1/response"
	"github.com/an-halim/golang-advanced/assignment-1/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	validator   *validator.Validate
}

// NewUserHandler membuat instance baru dari UserHandler
func NewUserHandler(userService service.IUserService, validator *validator.Validate) IUserHandler {
	return &UserHandler{
		userService: userService,
		validator:   validator,
	}
}

// CreateUser menghandle permintaan untuk membuat user baru
func (h *UserHandler) CreateUser(c *gin.Context) {
	request := request.CreateUserRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {

		errMsg := convertUserMandatoryFieldErrorString(err)
		apiResponse := response.Failed{}
		apiResponse.Status = "failed"
		apiResponse.Message = errMsg

		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	if err := h.validator.Struct(request); err != nil {
		apiResponse := response.Failed{}
		apiResponse.Status = "failed"
		apiResponse.Message = err.Error()

		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	userID, err := h.userService.CreateUser(c.Request.Context(), request)
	if err != nil {
		apiResponse := response.Failed{}
		apiResponse.Status = "failed"
		apiResponse.Message = err.Error()

		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	response := response.Success{}
	response.Message = fmt.Sprintf("User ID %d created successfully", userID)

	c.JSON(http.StatusCreated, response)
}

// GetUser menghandle permintaan untuk mendapatkan user berdasarkan ID
func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		apiResponse := response.Failed{}
		apiResponse.Status = "failed"
		apiResponse.Message = "Invalid ID"

		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	data, err := h.userService.GetUserByID(c.Request.Context(), id)
	if err != nil {
		fmt.Println(err)
		apiResponse := response.Failed{}
		apiResponse.Status = "failed"
		apiResponse.Message = "User not found"

		c.JSON(http.StatusNotFound, apiResponse)
		return
	}

	apiResponse := response.GetOne{}
	apiResponse.ID = data.ID
	apiResponse.Name = data.Name
	apiResponse.Email = data.Email
	apiResponse.RiskScore = data.RiskScore
	apiResponse.RiskCategory = data.RiskCategory
	apiResponse.RiskDefinition = data.RiskDefinition
	apiResponse.CreatedAt = data.CreatedAt
	apiResponse.UpdatedAt = data.UpdatedAt

	c.JSON(http.StatusOK, apiResponse)
}

// UpdateUser menghandle permintaan untuk mengupdate informasi user
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {

		apiResponse := response.Failed{}
		apiResponse.Status = "failed"
		apiResponse.Message = "Invalid ID"

		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	var request request.UpdateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {

		errMsg := convertUserMandatoryFieldErrorString(err)
		apiResponse := response.Failed{}
		apiResponse.Status = "failed"
		apiResponse.Message = errMsg

		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	if err := h.validator.Struct(request); err != nil {
		apiResponse := response.Failed{}
		apiResponse.Status = "failed"
		apiResponse.Message = err.Error()

		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	userID, err := h.userService.UpdateUser(c.Request.Context(), id, request)
	if err != nil {
		apiResponse := response.Failed{
			Status:  "failed",
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, apiResponse)
		return
	}

	response := response.Success{}
	response.Message = fmt.Sprintf("User ID %d updated successfully", userID)

	c.JSON(http.StatusOK, response)
}

// DeleteUser menghandle permintaan untuk menghapus user
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		apiResponse := response.Failed{}
		apiResponse.Status = "failed"
		apiResponse.Message = "Invalid ID"
		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	if err := h.userService.DeleteUser(c.Request.Context(), id); err != nil {
		apiResponse := response.Failed{}
		apiResponse.Status = "failed"
		apiResponse.Message = err.Error()
		c.JSON(http.StatusNotFound, apiResponse)
		return
	}

	apiResponse := response.Success{}
	apiResponse.Message = fmt.Sprintf("User ID %d deleted successfully", id)

	c.JSON(http.StatusOK, apiResponse)
}

// GetAllUsers menghandle permintaan untuk mendapatkan semua user
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	log.Println("Getting all users")
	pageSize := c.MustGet("pageSize").(int)
	currentPage := c.MustGet("currentPage").(int)
	offset := c.MustGet("offset").(int)

	users, err := h.userService.GetAllUsers(c.Request.Context(), pageSize, offset)
	if err != nil {
		apiResponse := response.Failed{}
		apiResponse.Status = "failed"
		apiResponse.Message = err.Error()
		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	apiResponse := response.ApiResponseGetAllUsers{}
	apiResponse.Page = currentPage
	apiResponse.Limit = pageSize
	apiResponse.Users = users

	c.JSON(http.StatusOK, apiResponse)
}

func convertUserMandatoryFieldErrorString(error error) string {
	switch {
	case strings.Contains(error.Error(), "'Name' failed on the 'required' tag"):
		return "name is mandatory"
	case strings.Contains(error.Error(), "'Email' failed on the 'required' tag"):
		return "email is mandatory"
	}
	return error.Error()
}
