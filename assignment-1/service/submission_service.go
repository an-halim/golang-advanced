package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/an-halim/golang-advanced/assignment-1/entity"
	"github.com/an-halim/golang-advanced/assignment-1/request"
	"github.com/an-halim/golang-advanced/assignment-1/response"
	"github.com/an-halim/golang-advanced/assignment-1/utils"
)

// IUserService mendefinisikan interface untuk layanan submission
type ISubmissionService interface {
	CreateSubmission(ctx context.Context, request request.CreateSubmissionInfo) (id int, error error)
	GetSubmissionByID(ctx context.Context, id int) (response.GetOneSubmission, error)
	UpdateSubmission(ctx context.Context, id int, submission entity.Submission) (entity.Submission, error)
	DeleteSubmission(ctx context.Context, id int) error
	GetAllSubmission(ctx context.Context, pageSize int, page int) ([]entity.Submission, error)
	GetAllUserSubmission(ctx context.Context, id int, pageSize int, page int) ([]entity.Submission, error)
}

// IUserRepository mendefinisikan interface untuk repository submission
type ISubmissionRepository interface {
	CreateSubmission(ctx context.Context, submission *entity.Submission) (entity.Submission, error)
	GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error)
	UpdateSubmission(ctx context.Context, id int, submission entity.Submission) (entity.Submission, error)
	DeleteSubmission(ctx context.Context, id int) error
	GetAllSubmission(ctx context.Context, pageSize int, page int) ([]entity.Submission, error)
	GetAllUserSubmission(ctx context.Context, id int, pageSize int, page int) ([]entity.Submission, error)
}

// submissionService adalah implementasi dari ISubmissionService yang menggunakan IUserRepository
type submissionService struct {
	submissionRepo ISubmissionRepository
	userService    IUserService
}

// NewSubmissionService membuat instance baru dari userService
func NewSubmissionService(submissionRepo ISubmissionRepository, userSuserService IUserService) ISubmissionService {
	return &submissionService{submissionRepo: submissionRepo, userService: userSuserService}
}

// CreateSubmission membuat submission baru
func (s *submissionService) CreateSubmission(ctx context.Context, input request.CreateSubmissionInfo) (id int, error error) {
	// Memanggil CreateSubmission dari repository untuk membuat submission baru
	// Convert the answers to JSON
	answersJSON, err := json.Marshal(input.Answers)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal answers: %v", err)
	}

	data := request.ValidateProfile(input.Answers)
	var submission entity.Submission
	submission.UserID = input.UserId
	submission.Answers = answersJSON
	submission.RiskCategory = string(data.Category)
	submission.RiskScore = data.MaxScore + data.MinScore

	user, err := s.userService.GetUserByID(ctx, input.UserId)
	if err != nil || user.ID == 0 {
		return 0, err
	}

	createSubmission, err := s.submissionRepo.CreateSubmission(ctx, &submission)
	if err != nil {
		return 0, fmt.Errorf("gagal membuat submission: %v", err)
	}
	return createSubmission.ID, nil
}

// GetUserByID mendapatkan submission berdasarkan ID
func (s *submissionService) GetSubmissionByID(ctx context.Context, id int) (response.GetOneSubmission, error) {
	// Memanggil GetSubmissionByID dari repository untuk mendapatkan submission berdasarkan ID
	submission, err := s.submissionRepo.GetSubmissionByID(ctx, id)
	if err != nil {
		return response.GetOneSubmission{}, fmt.Errorf("gagal mendapatkan submission berdasarkan ID: %v", err)
	}

	response := response.GetOneSubmission{}
	response.ID = submission.ID
	response.UserID = submission.User.ID
	response.Name = submission.User.Name
	response.Email = submission.User.Email
	response.Answer = submission.Answers
	response.RiskDefinition = utils.GetRiskProfileDefinition(submission.RiskScore)
	response.RiskCategory = submission.RiskCategory
	response.RiskScore = submission.RiskScore
	response.CreatedAt = submission.CreatedAt
	response.UpdatedAt = submission.UpdatedAt

	return response, nil
}

// UpdateUser memperbarui data submission
func (s *submissionService) UpdateSubmission(ctx context.Context, id int, submission entity.Submission) (entity.Submission, error) {
	// Memanggil UpdateUser dari repository untuk memperbarui data submission
	updatedSubmission, err := s.submissionRepo.UpdateSubmission(ctx, id, submission)
	if err != nil {
		return entity.Submission{}, fmt.Errorf("gagal memperbarui submission: %v", err)
	}
	return updatedSubmission, nil
}

// DeleteUser menghapus submission berdasarkan ID
func (s *submissionService) DeleteSubmission(ctx context.Context, id int) error {
	// Memanggil DeleteUser dari repository untuk menghapus submission berdasarkan ID
	err := s.submissionRepo.DeleteSubmission(ctx, id)
	if err != nil {
		return fmt.Errorf("gagal menghapus submission: %v", err)
	}
	return nil
}

// GetAllUsers mendapatkan semua submission
func (s *submissionService) GetAllSubmission(ctx context.Context, pageSize int, page int) ([]entity.Submission, error) {
	// Memanggil GetAllUsers dari repository untuk mendapatkan semua submission
	submissions, err := s.submissionRepo.GetAllSubmission(ctx, pageSize, page)
	if err != nil {
		return nil, fmt.Errorf("gagal mendapatkan semua submission: %v", err)
	}
	return submissions, nil
}

// GetAllUsers mendapatkan semua submission
func (s *submissionService) GetAllUserSubmission(ctx context.Context, id int, pageSize int, page int) ([]entity.Submission, error) {
	// Memanggil GetAllUsers dari repository untuk mendapatkan semua submission
	submissions, err := s.submissionRepo.GetAllUserSubmission(ctx, id, pageSize, page)
	if err != nil {
		return nil, fmt.Errorf("gagal mendapatkan semua submission: %v", err)
	}
	return submissions, nil
}
