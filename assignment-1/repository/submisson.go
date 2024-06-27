package repository

import (
	"context"
	"errors"
	"log"

	"github.com/an-halim/golang-advanced/assignment-1/entity"
	"github.com/an-halim/golang-advanced/assignment-1/service"
	"gorm.io/gorm"
)

type submissionRepository struct {
	db GormDBIface
}

// NewSubmissionRepository membuat instance baru dari submissionRepository
func NewSubmissionRepository(db GormDBIface) service.ISubmissionRepository {
	return &submissionRepository{db: db}
}

// CreateSubmission membuat submission baru dalam basis data
func (r *submissionRepository) CreateSubmission(ctx context.Context, submission *entity.Submission) (entity.Submission, error) {
	if err := r.db.WithContext(ctx).Create(submission).Error; err != nil {
		log.Printf("Error creating submission: %v\n", err)
		return entity.Submission{}, err
	}
	return *submission, nil
}

// GetSubmissionByID mengambil submission berdasarkan ID
func (r *submissionRepository) GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error) {
	var submission entity.Submission
	if err := r.db.WithContext(ctx).Preload("User").First(&submission, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Submission{}, nil
		}
		log.Printf("Error getting submission by ID: %v\n", err)
		return entity.Submission{}, err
	}
	return submission, nil
}

// UpdateSubmission memperbarui informasi submission dalam basis data
func (r *submissionRepository) UpdateSubmission(ctx context.Context, id int, submission entity.Submission) (entity.Submission, error) {
	return entity.Submission{}, nil
}

// Deletesubmission menghapus submission berdasarkan ID
func (r *submissionRepository) DeleteSubmission(ctx context.Context, id int) error {

	if err := r.db.WithContext(ctx).Delete(&entity.Submission{}, id).Error; err != nil {
		log.Printf("Error deleting submission: %v\n", err)
		return err
	}
	return nil
}

// GetAllSubmission mengambil semua submission dari basis data
func (r *submissionRepository) GetAllSubmission(ctx context.Context, pageSize int, offset int) ([]entity.Submission, error) {
	log.Println("Getting all submissions")
	var submissions []entity.Submission
	if err := r.db.WithContext(ctx).Limit(pageSize).Offset(offset).Find(&submissions).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return submissions, nil
		}
		log.Printf("Error getting all submissions: %v\n", err)
		return nil, err
	}
	return submissions, nil
}

// GetAllSubmission mengambil semua submission dari basis data
func (r *submissionRepository) GetAllUserSubmission(ctx context.Context, id int, pageSize int, offset int) ([]entity.Submission, error) {
	log.Printf("Getting all submissions for user with ID: %d\n", id)
	var submissions []entity.Submission
	if err := r.db.WithContext(ctx).Where("user_id = ?", id).Limit(pageSize).Offset(offset).Find(&submissions).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return submissions, nil
		}
		log.Printf("Error getting all submissions: %v\n", err)
		return nil, err
	}
	return submissions, nil
}
