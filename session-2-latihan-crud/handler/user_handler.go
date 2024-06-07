package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/an-halim/golang-advanced/session-2-latihan-crud/entity"
	"github.com/an-halim/golang-advanced/session-2-latihan-crud/response"
	"github.com/gin-gonic/gin"
)

var (
	users  []entity.User
	nextID int
)

func CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		apiResponse := response.APIResponseFailed{}
		apiResponse.Status = "failed"
		apiResponse.Message = err.Error()

		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	user.ID = nextID
	nextID++
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	users = append(users, user)

	apiResponse := response.APIResponseSuccess{}
	apiResponse.Status = "success"
	apiResponse.Message = "User created"
	apiResponse.Data = user

	c.JSON(http.StatusCreated, apiResponse)
}

func GetUsers(c *gin.Context) {
	// pagination
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		apiResponse := response.APIResponseFailed{}
		apiResponse.Status = "failed"
		apiResponse.Message = err.Error()

		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		apiResponse := response.APIResponseFailed{}
		apiResponse.Status = "failed"
		apiResponse.Message = err.Error()

		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	// calculate the start and end index
	startIndex := (page - 1) * limit
	endIndex := page * limit

	if endIndex > len(users) {
		endIndex = len(users)
	}

	// use slice to get the data based on pagination
	apiResponse := response.APIResponseGetAll{}
	apiResponse.Status = "success"
	apiResponse.Message = "List of users"
	apiResponse.Page = page
	apiResponse.Limit = limit
	apiResponse.Data = users[startIndex:endIndex]

	c.JSON(http.StatusOK, apiResponse)
}

func GetUserbyID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		apiResponse := response.APIResponseFailed{}
		apiResponse.Status = "failed"
		apiResponse.Message = "Invalid ID"

		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	for _, user := range users {
		if user.ID == id {
			apiResponse := response.APIResponseSuccess{}
			apiResponse.Status = "success"
			apiResponse.Message = "User found"
			apiResponse.Data = user

			c.JSON(http.StatusOK, apiResponse)
			return
		}
	}

	apiResponse := response.APIResponseFailed{}
	apiResponse.Status = "failed"
	apiResponse.Message = "User not found"

	c.JSON(http.StatusBadRequest, apiResponse)
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		apiResponse := response.APIResponseFailed{}
		apiResponse.Status = "failed"
		apiResponse.Message = "Invalid ID"

		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		apiResponse := response.APIResponseFailed{}
		apiResponse.Status = "failed"
		apiResponse.Message = err.Error()

		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	for i, u := range users {
		if u.ID == id {
			newUser := entity.User{
				ID:        id,
				Name:      user.Name,
				Email:     user.Email,
				Password:  u.Password,
				CreatedAt: u.CreatedAt,
				UpdatedAt: time.Now(),
			}

			users[i] = newUser
			apiResponse := response.APIResponseSuccess{}
			apiResponse.Status = "success"
			apiResponse.Message = "User updated"
			apiResponse.Data = user

			c.JSON(http.StatusOK, apiResponse)
			return
		}
	}

	apiResponse := response.APIResponseFailed{}
	apiResponse.Status = "failed"
	apiResponse.Message = "User not found"

	c.JSON(http.StatusBadRequest, apiResponse)
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		apiResponse := response.APIResponseFailed{}
		apiResponse.Status = "failed"
		apiResponse.Message = "Invalid ID"

		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			apiResponse := response.APIResponseSuccess{}
			apiResponse.Status = "success"
			apiResponse.Message = "User deleted"

			c.JSON(http.StatusOK, apiResponse)
			return
		}
	}

	apiResponse := response.APIResponseFailed{}
	apiResponse.Status = "failed"
	apiResponse.Message = "User not found"

	c.JSON(http.StatusBadRequest, apiResponse)
}
