package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vctrthe/api-go/auth"
	"github.com/vctrthe/api-go/campaign"
	"github.com/vctrthe/api-go/config"
	"github.com/vctrthe/api-go/payment"
	"github.com/vctrthe/api-go/routes"
	"github.com/vctrthe/api-go/transaction"
	"github.com/vctrthe/api-go/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Initialize the configuration
	config.ReadConfig()

	// Initialize the database using the values from the config
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		config.C.Database.Username,
		config.C.Database.Password,
		config.C.Database.Host,
		config.C.Database.Port,
		config.C.Database.DBName,
		config.C.Database.Charset,
		config.C.Database.ParseTime,
		config.C.Database.Loc,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	// Repositories
	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)
	// Services
	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	paymentService := payment.NewService()
	transactionService := transaction.NewService(transactionRepository, campaignRepository, paymentService)
	authService, err := auth.NewService(config.C.JWT.Secret)
	if err != nil {
		log.Fatal("Error initializing JWT Service", err)
	}

	// API Endpoint Routes
	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/images", "./images")

	routes.RegisterRoutes(router, authService, userService, campaignService, transactionService, paymentService)
	router.Run()
}
