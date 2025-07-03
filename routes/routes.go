package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wahyuutomoputra/order-management/handler"
	"github.com/wahyuutomoputra/order-management/middleware"
	"github.com/wahyuutomoputra/order-management/repository"
	"github.com/wahyuutomoputra/order-management/service"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Dependency injection
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	authHandler := handler.NewAuthHandler(userService)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	r.POST("/register", authHandler.RegisterHandler())
	r.POST("/login", authHandler.LoginHandler())
	r.GET("/me", middleware.AuthMiddleware(), authHandler.MeHandler())

	// Produk
	product := r.Group("/products")
	{
		product.GET("", productHandler.ListProductHandler())
		product.GET(":id", productHandler.GetProductHandler())
	}

	admin := r.Group("/admin", middleware.AuthMiddleware(), handler.AdminOnly())
	{
		admin.POST("/products", productHandler.CreateProductHandler())
		admin.PUT("/products/:id", productHandler.UpdateProductHandler())
		admin.DELETE("/products/:id", productHandler.DeleteProductHandler())
	}

	r.POST("/orders", middleware.AuthMiddleware(), orderHandler.CreateOrderHandler())
	r.GET("/orders/history", middleware.AuthMiddleware(), orderHandler.OrderHistoryHandler())
}
