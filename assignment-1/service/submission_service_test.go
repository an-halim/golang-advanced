package service_test

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/an-halim/golang-advanced/assignment-1/entity"
	"github.com/an-halim/golang-advanced/assignment-1/request"
	"github.com/an-halim/golang-advanced/assignment-1/service"
	mock_service "github.com/an-halim/golang-advanced/assignment-1/test/mock/service"
	"github.com/golang/mock/gomock"

	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSubmissionService_CreateSubmission(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_service.NewMockISubmissionRepository(ctrl)
	userMock := mock_service.NewMockIUserRepository(ctrl)
	submissionService := service.NewSubmissionService(mockRepo, userMock)

	ctx := context.Background()
	submission := &entity.Submission{
		UserID:       1,
		RiskScore:    50,
		RiskCategory: "Medium",
		Answers:      json.RawMessage(`{"question_id":1, "answer":"answer1"}`),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		User:         entity.User{ID: 1, Name: "User", Email: "lorem@gmail.com"},
	}

	var answers []request.Answers
	_ = json.Unmarshal(submission.Answers, &answers)

	request := request.CreateSubmissionInfo{
		UserId:  submission.ID,
		Answers: answers,
	}
	t.Run("PositiveCase", func(t *testing.T) {
		mockRepo.EXPECT().CreateSubmission(ctx, submission).Return(*submission, nil)

		createdSubmission, err := submissionService.CreateSubmission(ctx, request)
		assert.NoError(t, err)
		assert.Equal(t, *submission, createdSubmission)
	})

	t.Run("NegativeCase", func(t *testing.T) {
		mockRepo.EXPECT().CreateSubmission(ctx, submission).Return(entity.Submission{}, errors.New("failed to create user"))

		_, err := submissionService.CreateSubmission(ctx, request)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create user")
	})
}
