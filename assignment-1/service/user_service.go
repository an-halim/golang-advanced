package service

import (
	"context"
	"fmt"

	"github.com/an-halim/golang-advanced/assignment-1/entity"
	"github.com/an-halim/golang-advanced/assignment-1/request"
	"github.com/an-halim/golang-advanced/assignment-1/response"
	"github.com/an-halim/golang-advanced/assignment-1/utils"
)

// IUserService mendefinisikan interface untuk layanan pengguna
type IUserService interface {
	CreateUser(ctx context.Context, request request.CreateUserRequest) (resultId int, error error)
	GetUserByID(ctx context.Context, id int) (data response.ResponseSubmissionInfo, err error)
	UpdateUser(ctx context.Context, id int, request request.UpdateUserRequest) (resultId int, error error)
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context, pageSize int, page int) (data []response.ResponseSubmissionInfo, error error)
}

// IUserRepository mendefinisikan interface untuk repository pengguna
type IUserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context, pageSize int, page int) ([]entity.User, error)
}

// userService adalah implementasi dari IUserService yang menggunakan IUserRepository
type userService struct {
	userRepo IUserRepository
}

// NewUserService membuat instance baru dari userService
func NewUserService(userRepo IUserRepository) IUserService {
	return &userService{userRepo: userRepo}
}

// CreateUser membuat pengguna baru
func (s *userService) CreateUser(ctx context.Context, request request.CreateUserRequest) (resultId int, error error) {
	// Memanggil CreateUser dari repository untuk membuat pengguna baru
	user := new(entity.User)
	user.Name = request.Name
	user.Email = request.Email

	result, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("gagal membuat pengguna: %v", err)
	}
	return result.ID, nil
}

// GetUserByID mendapatkan pengguna berdasarkan ID
func (s *userService) GetUserByID(ctx context.Context, id int) (response response.ResponseSubmissionInfo, err error) {
	// Memanggil GetUserByID dari repository untuk mendapatkan pengguna berdasarkan ID
	result, err := s.userRepo.GetUserByID(ctx, id)

	if err != nil || result.ID == 0 {
		return response, fmt.Errorf("failed to get user")
	}

	response.ID = result.ID
	response.Name = result.Name
	response.Email = result.Email
	response.CreatedAt = result.CreatedAt
	response.UpdatedAt = result.UpdatedAt
	if len(result.Submissions) != 0 {
		response.RiskCategory = result.Submissions[0].RiskCategory
		response.RiskScore = result.Submissions[0].RiskScore
		response.RiskDefinition = utils.GetRiskProfileDefinition(response.RiskScore)
	}

	return response, nil
}

// UpdateUser memperbarui data pengguna
func (s *userService) UpdateUser(ctx context.Context, id int, request request.UpdateUserRequest) (resultId int, error error) {
	// Memanggil UpdateUser dari repository untuk memperbarui data pengguna
	user := entity.User{
		Name:  request.Name,
		Email: request.Email,
	}
	updatedUser, err := s.userRepo.UpdateUser(ctx, id, user)
	if err != nil {
		return 0, fmt.Errorf("gagal memperbarui pengguna: %v", err)
	}
	return updatedUser.ID, nil
}

// DeleteUser menghapus pengguna berdasarkan ID
func (s *userService) DeleteUser(ctx context.Context, id int) error {
	// Memanggil DeleteUser dari repository untuk menghapus pengguna berdasarkan ID
	user, errU := s.userRepo.GetUserByID(ctx, id)
	if errU != nil || user.ID == 0 {
		return fmt.Errorf("failed to delete user")
	}

	err := s.userRepo.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("gagal menghapus pengguna: %v", err)
	}
	return nil
}

// GetAllUsers mendapatkan semua pengguna
func (s *userService) GetAllUsers(ctx context.Context, pageSize int, page int) (data []response.ResponseSubmissionInfo, error error) {
	// Memanggil GetAllUsers dari repository untuk mendapatkan semua pengguna
	result, err := s.userRepo.GetAllUsers(ctx, pageSize, page)

	for _, user := range result {
		var response response.ResponseSubmissionInfo
		response.ID = user.ID
		response.Name = user.Name
		response.Email = user.Email
		if len(user.Submissions) == 0 {
			response.RiskCategory = ""
			response.RiskScore = 0
			response.RiskDefinition = ""
		} else {
			response.RiskCategory = user.Submissions[0].RiskCategory
			response.RiskScore = user.Submissions[0].RiskScore
			response.RiskDefinition = utils.GetRiskProfileDefinition(response.RiskScore)
		}
		response.CreatedAt = user.CreatedAt
		response.UpdatedAt = user.UpdatedAt

		data = append(data, response)
	}

	if err != nil {
		return data, fmt.Errorf("gagal mendapatkan semua pengguna: %v", err)
	}
	return data, nil
}
