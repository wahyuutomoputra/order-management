package main

import (
	"fmt"
	"log"
	"os"

	"github.com/wahyuutomoputra/order-management/config"
	"github.com/wahyuutomoputra/order-management/models"
	"github.com/wahyuutomoputra/order-management/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/wahyuutomoputra/order-management/docs"
	"golang.org/x/crypto/bcrypt"
)

// @title Order Management API
// @version 1.0
// @description API for order management system
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	fmt.Println("Starting Order Management API...")
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	log.Println("Connected to database.")
	// Auto migrate
	if err := db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.OrderItem{}); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	} else {
		log.Println("Database migration successful")
	}

	var count int64
	db.Model(&models.User{}).Where("role = ?", "admin").Count(&count)
	if count == 0 {
		hash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		db.Create(&models.User{
			Name:     "Super Admin",
			Email:    "admin@gmail.com",
			Password: string(hash),
			Role:     "admin",
		})
		log.Println("Default admin created: admin@gmail.com / admin123")
	}

	r := gin.Default()

	routes.SetupRoutes(r, db)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
