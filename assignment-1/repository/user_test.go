package repository

// import (
// 	"context"
// 	"reflect"
// 	"testing"

// 	"github.com/stretchr/testify/assert" // Requires testify, install with: go get github.com/stretchr/testify/assert
// 	"github.com/stretchr/testify/mock"
// 	"github.com/your-username/your-project/entity" // Update with your actual package path
// 	"gorm.io/gorm"
// )

// // Mock Repository (Assuming you're using an interface for your repository)
// type mockUserRepository struct {
// 	mock.Mock
// }

// func (m *mockUserRepository) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
// 	args := m.Called(ctx, user)
// 	return args.Get(0).(entity.User), args.Error(1)
// }

// func Test_userRepository_CreateUser(t *testing.T) {
// 	type args struct {
// 		ctx  context.Context
// 		user *entity.User
// 	}
// 	tests := []struct {
// 		name    string
// 		r       *userRepository
// 		args    args
// 		want    entity.User
// 		wantErr bool
// 	}{
// 		{
// 			name: "CreateUser_Success",
// 			r: &userRepository{
// 				// Inject your mock repository here
// 				UserRepository: &mockUserRepository{},
// 			},
// 			args: args{
// 				ctx: context.Background(),
// 				user: &entity.User{
// 					// Provide test user data
// 					// Example:
// 					ID:    1,
// 					Name:  "Test User",
// 					Email: "test@example.com",
// 				},
// 			},
// 			want: entity.User{
// 				// Expected user data after creation
// 				// Example:
// 				ID:    1,
// 				Name:  "Test User",
// 				Email: "test@example.com",
// 			},
// 			wantErr: false,
// 		},
// 		// Add more test cases for different scenarios:
// 		// - Error handling (e.g., database errors)
// 		// - Validation errors (e.g., invalid email format)
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Type assertion to access the mock methods
// 			mockRepo := tt.r.UserRepository.(*mockUserRepository)

// 			// Set up expectations for the mock repository
// 			mockRepo.On("CreateUser", tt.args.ctx, tt.args.user).Return(tt.want, nil)

// 			got, err := tt.r.CreateUser(tt.args.ctx, tt.args.user)

// 			// Assertions using testify
// 			assert.Equal(t, tt.wantErr, err != nil, "Error expectation mismatch")
// 			assert.Equal(t, tt.want, got, "Created user data mismatch")

// 			// Assert that the mock method was called as expected
// 			mockRepo.AssertExpectations(t)
// 		})
// 	}
// }

// func Test_userRepository_GetUserByID(t *testing.T) {
// 	type args struct {
// 		ctx context.Context
// 		id  int
// 	}
// 	tests := []struct {
// 		name    string
// 		r       *userRepository
// 		args    args
// 		want    entity.User
// 		wantErr bool
// 	}{
// 		{
// 			name: "Test case 1",
// 			r: &userRepository{
// 				db: &gorm.DB{},
// 			},
// 			args: args{
// 				ctx: context.Background(),
// 				id:  1,
// 			},
// 			want:    entity.User{},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := tt.r.GetUserByID(tt.args.ctx, tt.args.id)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("userRepository.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("userRepository.GetUserByID() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
