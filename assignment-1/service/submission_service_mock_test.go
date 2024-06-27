package service_test

import (
	"time"

	"github.com/an-halim/golang-advanced/assignment-1/entity"
)

// MockSubmissionRepository adalah mock untuk IUserRepository
type MockSubmissionRepository struct {
	submission []entity.Submission
}

func (m *MockSubmissionRepository) CreateSubmission(submission *entity.Submission) entity.Submission {
	submission.ID = len(m.submission)
	submission.UserID = 1
	submission.RiskCategory = "Low"
	submission.RiskScore = 0
	submission.Answers = []byte(`{"answers": "answers"}`)
	submission.User = entity.User{ID: 1, Name: "User", Email: "lorem@gmail.com"}
	submission.CreatedAt = time.Now()
	submission.UpdatedAt = time.Now()
	m.submission = append(m.submission, *submission)
	return *submission
}

func (m *MockSubmissionRepository) GetSubmissionByID(id int) (entity.Submission, bool) {
	for _, submission := range m.submission {
		if submission.ID == id {
			return submission, true
		}
	}
	return entity.Submission{}, false
}

func (m *MockSubmissionRepository) UpdateSubmission(id int, submission entity.Submission) (entity.Submission, bool) {
	for i, s := range m.submission {
		if s.ID == id {
			submission.ID = id
			submission.UserID = id
			submission.RiskCategory = "Medium"
			submission.RiskScore = 10
			submission.Answers = []byte(`{"answers": "answers"}`)
			submission.User = entity.User{ID: 1, Name: "User", Email: "lorem@gmail.com"}
			submission.CreatedAt = time.Now()
			submission.UpdatedAt = time.Now()
			m.submission[i] = submission
			return submission, true
		}
	}
	return entity.Submission{}, false
}

func (m *MockSubmissionRepository) DeleteSubmission(id int) bool {
	for i, user := range m.submission {
		if user.ID == id {
			m.submission = append(m.submission[:i], m.submission[i+1:]...)
			return true
		}
	}
	return false
}

func (m *MockSubmissionRepository) GetAllSubmission() []entity.Submission {
	return m.submission
}
