package repository_test

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/an-halim/golang-advanced/assignment-1/entity"
	"github.com/an-halim/golang-advanced/assignment-1/repository"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupSQLMockSubmission(t *testing.T) (sqlmock.Sqlmock, *gorm.DB) {
	// Setup SQL mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// Setup GORM with the mock DB
	gormDB, gormDBErr := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if gormDBErr != nil {
		t.Fatalf("failed to open GORM connection: %v", gormDBErr)
	}
	return mock, gormDB
}

func TestSubmissionRepository_CreateSubmission(t *testing.T) {
	mock, gormDB := setupSQLMockSubmission(t)
	submissionRepo := repository.NewSubmissionRepository(gormDB)

	expectedQueryString := regexp.QuoteMeta(`INSERT INTO "submissions" ("user_id","risk_score","risk_category","answers","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)

	t.Run("Positive Case", func(t *testing.T) {
		submission := &entity.Submission{
			UserID:       1,
			RiskScore:    50,
			RiskCategory: "Medium",
			Answers:      json.RawMessage(`{"question":"answer1"}`),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		mock.ExpectQuery(expectedQueryString).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).
				AddRow(1))

		createdSubmission, err := submissionRepo.CreateSubmission(context.Background(), submission)
		require.NoError(t, err)
		require.NotNil(t, createdSubmission.ID)
		require.Equal(t, submission.UserID, createdSubmission.UserID)
		require.Equal(t, submission.RiskScore, createdSubmission.RiskScore)
		require.Equal(t, submission.RiskCategory, createdSubmission.RiskCategory)
	})

	t.Run("Negative Case", func(t *testing.T) {
		submission := &entity.Submission{
			UserID:       1,
			RiskScore:    50,
			RiskCategory: "Medium",
			Answers:      json.RawMessage(`{"question1":"answer1"}`),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		mock.ExpectQuery(expectedQueryString).
			WillReturnError(errors.New("db error"))

		createdSubmission, err := submissionRepo.CreateSubmission(context.Background(), submission)
		require.Error(t, err)
		require.Empty(t, createdSubmission)
	})
}

func TestSubmissionRepository_GetSubmissionByID(t *testing.T) {
	mock, gormDB := setupSQLMockSubmission(t)
	submissionRepo := repository.NewSubmissionRepository(gormDB)

	expectedSubmissionQueryString := regexp.QuoteMeta(`SELECT * FROM "submissions" WHERE "submissions"."id" = $1 ORDER BY "submissions"."id" LIMIT $2`)
	expectedUserQueryString := regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1`)

	t.Run("Positive Case", func(t *testing.T) {
		mock.ExpectQuery(expectedSubmissionQueryString).
			WithArgs(1, 1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "risk_score", "risk_category", "answers", "created_at", "updated_at"}).
				AddRow(1, 1, 50, "Medium", json.RawMessage(`{"question1":"answer1"}`), time.Now(), time.Now()))

		mock.ExpectQuery(expectedUserQueryString).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
				AddRow(1, "John Doe"))

		submission, err := submissionRepo.GetSubmissionByID(context.Background(), 1)
		expectedAnswer := json.RawMessage(`{"question1":"answer1"}`)
		require.NoError(t, err)
		require.Equal(t, 1, submission.ID)
		require.Equal(t, 1, submission.UserID)
		require.Equal(t, 50, submission.RiskScore)
		require.Equal(t, "Medium", submission.RiskCategory)
		require.Equal(t, expectedAnswer, submission.Answers)
		require.Equal(t, 1, submission.User.ID)
		require.Equal(t, "John Doe", submission.User.Name)
	})

	t.Run("No data found Case", func(t *testing.T) {
		mock.ExpectQuery(expectedSubmissionQueryString).
			WithArgs(1, 1).
			WillReturnError(gorm.ErrRecordNotFound)

		submission, err := submissionRepo.GetSubmissionByID(context.Background(), 1)
		require.NoError(t, err)
		require.Empty(t, submission)
	})

	t.Run("Negative Case", func(t *testing.T) {
		mock.ExpectQuery(expectedSubmissionQueryString).
			WithArgs(1, 1).
			WillReturnError(errors.New("db error"))

		submission, err := submissionRepo.GetSubmissionByID(context.Background(), 1)
		require.Error(t, err)
		require.Empty(t, submission)
	})
}

func TestSubmissionRepository_UpdateSubmissions(t *testing.T) {
	mock, gormDB := setupSQLMockSubmission(t)
	submissionRepo := repository.NewSubmissionRepository(gormDB)

	expectedUpdateSubmissionQueryString := regexp.QuoteMeta(`UPDATE "submissions" SET "user_id"=$1, "risk_score"=$2, "risk_category"=$3, "answers"=$4, "created_at"=$5, "updated_at"=$6 WHERE "id" = $7`)
	expectedSelectSubmissionQueryString := regexp.QuoteMeta(`SELECT "id", "user_id", "risk_score", "risk_category", "answers", "created_at", "updated_at" FROM "submissions" WHERE "id" = $1`)

	t.Run("Update Submission Case", func(t *testing.T) {
		submission := entity.Submission{
			ID:           1,
			UserID:       1,
			RiskScore:    70,
			RiskCategory: "High",
			Answers:      json.RawMessage(`{"question1":"answer1_updated"}`),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		mock.ExpectExec(expectedUpdateSubmissionQueryString).
			WithArgs(submission.UserID, submission.RiskScore, submission.RiskCategory, submission.Answers, submission.CreatedAt, submission.UpdatedAt, submission.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectQuery(expectedSelectSubmissionQueryString).
			WithArgs(submission.ID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "risk_score", "risk_category", "answers", "created_at", "updated_at"}).
				AddRow(submission.ID, submission.UserID, submission.RiskScore, submission.RiskCategory, submission.Answers, submission.CreatedAt, submission.UpdatedAt))

		data, err := submissionRepo.UpdateSubmission(context.Background(), submission.ID, submission)
		require.NoError(t, err)
		log.Println(data)
	})
}
