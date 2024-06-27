package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/an-halim/golang-advanced/assignment-1/request"
	"github.com/an-halim/golang-advanced/assignment-1/response"
	"github.com/an-halim/golang-advanced/assignment-1/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// IsubmissionHandler mendefinisikan interface untuk handler submission
type ISubmissionHandler interface {
	CreateSubmission(c *gin.Context)
	GetSubmissionByID(c *gin.Context)
	DeleteSubmission(c *gin.Context)
	GetAllSubmission(c *gin.Context)
}

type SubmissionHandler struct {
	submissionService service.ISubmissionService
	validator         *validator.Validate
}

// NewsubmissionHandler membuat instance baru dari submissionHandler
func NewSubmissionHandler(submissionService service.ISubmissionService, validator *validator.Validate) ISubmissionHandler {
	return &SubmissionHandler{
		submissionService: submissionService,
		validator:         validator,
	}
}

// Createsubmission menghandle permintaan untuk membuat submission baru
func (h *SubmissionHandler) CreateSubmission(c *gin.Context) {
	var requestInfo request.CreateSubmissionInfo
	if err := c.ShouldBindJSON(&requestInfo); err != nil {
		errMsg := err.Error()
		apiResponse := response.Failed{
			Status:  "failed",
			Message: errMsg,
		}
		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	if err := h.validator.Struct(requestInfo); err != nil {
		apiResponse := response.Failed{}
		apiResponse.Status = "failed"
		apiResponse.Message = err.Error()

		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	id, err := h.submissionService.CreateSubmission(c.Request.Context(), requestInfo)
	if err != nil {
		apiResponse := response.Failed{
			Status:  "failed",
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	apiResponse := response.Success{}
	apiResponse.Message = fmt.Sprintf("submission ID %d created successfully", id)

	c.JSON(http.StatusCreated, apiResponse)
}

// Getsubmission menghandle permintaan untuk mendapatkan submission berdasarkan ID
func (h *SubmissionHandler) GetSubmissionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	submission, err := h.submissionService.GetSubmissionByID(c.Request.Context(), id)
	if err != nil {
		apiResponse := response.Failed{
			Status:  "failed",
			Message: "submission not found",
		}
		c.JSON(http.StatusNotFound, apiResponse)
		return
	}

	c.JSON(http.StatusOK, submission)
}

// Deletesubmission menghandle permintaan untuk menghapus submission
func (h *SubmissionHandler) DeleteSubmission(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		apiResponse := response.Failed{
			Status:  "failed",
			Message: "Invalid ID",
		}
		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	if err := h.submissionService.DeleteSubmission(c.Request.Context(), id); err != nil {
		apiResponse := response.Failed{
			Status:  "failed",
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, apiResponse)
		return
	}

	apiResponse := response.Success{
		Message: "submission deleted successfully",
	}

	c.JSON(http.StatusOK, apiResponse)
}

// GetAllsubmissions menghandle permintaan untuk mendapatkan semua submission
func (h *SubmissionHandler) GetAllSubmission(c *gin.Context) {
	userId, _ := strconv.Atoi(c.DefaultQuery("user_id", "0"))
	pageSize := c.MustGet("pageSize").(int)
	currentPage := c.MustGet("currentPage").(int)
	offset := c.MustGet("offset").(int)

	if userId == 0 {
		submissions, err := h.submissionService.GetAllSubmission(c.Request.Context(), pageSize, offset)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response := response.ApiResponseGetAllSubmission{}
		response.Limit = offset
		response.Page = currentPage
		response.Submission = submissions

		c.JSON(http.StatusOK, response)
		return
	}

	submissions, err := h.submissionService.GetAllUserSubmission(c.Request.Context(), userId, pageSize, offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := response.ApiResponseGetAllSubmission{}
	response.UserID = userId
	response.Limit = pageSize
	response.Page = currentPage
	response.TotalSubmission = len(submissions)
	response.Submission = submissions

	c.JSON(http.StatusOK, response)
}
