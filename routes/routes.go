// routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vctrthe/api-go/auth"
	"github.com/vctrthe/api-go/campaign"
	"github.com/vctrthe/api-go/handler"
	"github.com/vctrthe/api-go/middleware"
	"github.com/vctrthe/api-go/payment"
	"github.com/vctrthe/api-go/transaction"
	"github.com/vctrthe/api-go/user"
)

// RegisterRoutes registers API routes and handlers
func RegisterRoutes(router *gin.Engine, authService auth.Service, userService user.Service, campaignService campaign.Service, transactionService transaction.Service, paymentService payment.Service) {
	api := router.Group("/api/v1")

	// User-related endpoints
	userHandler := handler.NewUserHandler(userService, authService)
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.LoginUser)
	api.POST("/email_check", userHandler.CheckEmail)
	api.POST("/avatars", middleware.AuthMiddleware(authService, userService), userHandler.AvatarUpload)

	// Campaign-related endpoints
	campaignHandler := handler.NewCampaignHandler(campaignService)
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", middleware.AuthMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", middleware.AuthMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", middleware.AuthMiddleware(authService, userService), campaignHandler.UploadImage)

	// Transaction-related endpoints
	transactionHandler := handler.NewTransactionHandler(transactionService)
	api.GET("/campaigns/:id/transactions", middleware.AuthMiddleware(authService, userService), transactionHandler.GetCampaignTransactions)
	api.GET("/transactions", middleware.AuthMiddleware(authService, userService), transactionHandler.GetUserTransaction)
	api.POST("/transactions", middleware.AuthMiddleware(authService, userService), transactionHandler.CreateTransaction)
	api.POST("/transactions/notification", transactionHandler.GetNotification)
}
