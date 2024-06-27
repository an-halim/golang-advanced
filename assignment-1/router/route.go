// package router mengatur rute untuk aplikasi
package router

import (
	"github.com/an-halim/golang-advanced/assignment-1/handler"
	"github.com/an-halim/golang-advanced/assignment-1/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRouter menginisialisasi dan mengatur rute untuk aplikasi
func SetupRouter(r *gin.Engine, userHandler handler.IUserHandler, submissionHandler handler.ISubmissionHandler) {
	// Mengatur endpoint publik untuk pengguna
	usersPublicEndpoint := r.Group("/users")
	usersPublicEndpoint.Use(middleware.PaginationMiddleware())
	// Rute untuk mendapatkan pengguna berdasarkan ID
	usersPublicEndpoint.GET("/:id", userHandler.GetUser)
	// Rute untuk mendapatkan semua pengguna
	usersPublicEndpoint.GET("", userHandler.GetAllUsers)
	usersPublicEndpoint.GET("/", userHandler.GetAllUsers)

	// Rute untuk membuat pengguna baru
	usersPublicEndpoint.POST("", userHandler.CreateUser)
	usersPublicEndpoint.POST("/", userHandler.CreateUser)
	// Rute untuk memperbarui pengguna berdasarkan ID
	usersPublicEndpoint.PUT("/:id", userHandler.UpdateUser)
	// Rute untuk menghapus pengguna berdasarkan ID
	usersPublicEndpoint.DELETE("/:id", userHandler.DeleteUser)

	// Mengatur endpoint publik untuk pengguna
	submissionPublicEndpoint := r.Group("/submissions")
	submissionPublicEndpoint.Use(middleware.PaginationMiddleware())
	// Rute untuk mendapatkan pengguna berdasarkan ID
	submissionPublicEndpoint.GET("/:id", submissionHandler.GetSubmissionByID)
	// Rute untuk mendapatkan semua pengguna
	submissionPublicEndpoint.GET("", submissionHandler.GetAllSubmission)
	submissionPublicEndpoint.GET("/", submissionHandler.GetAllSubmission)

	// Rute untuk membuat pengguna baru
	submissionPublicEndpoint.POST("", submissionHandler.CreateSubmission)
	submissionPublicEndpoint.POST("/", submissionHandler.CreateSubmission)
	// Rute untuk menghapus pengguna berdasarkan ID
	submissionPublicEndpoint.DELETE("/:id", submissionHandler.DeleteSubmission)
}
