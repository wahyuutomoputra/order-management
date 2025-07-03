package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/wahyuutomoputra/order-management/middleware"
	"github.com/wahyuutomoputra/order-management/models"
	"github.com/wahyuutomoputra/order-management/service"
	"github.com/wahyuutomoputra/order-management/utils"
)

var validate = validator.New()

type AuthHandler struct {
	UserService *service.UserService
}

func NewAuthHandler(userService *service.UserService) *AuthHandler {
	return &AuthHandler{UserService: userService}
}

// Register & Login Request struct

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// RegisterHandler godoc
// @Security BearerAuth
// @Summary Register user
// @Tags Auth
// @Accept json
// @Produce json
// @Param data body RegisterRequest true "Register data"
// @Success 201 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Router /register [post]
func (h *AuthHandler) RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.JSONError(c, 400, "Invalid request")
			return
		}
		if err := validate.Struct(req); err != nil {
			utils.JSONError(c, 400, err.Error())
			return
		}
		if _, err := h.UserService.FindByEmail(req.Email); err == nil {
			utils.JSONError(c, 400, "Email already registered")
			return
		}
		hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		user := models.User{
			Name:     req.Name,
			Email:    req.Email,
			Password: string(hash),
			Role:     "customer",
		}
		if err := h.UserService.Register(&user); err != nil {
			utils.JSONError(c, 500, "Failed to register user")
			return
		}
		utils.JSONCreated(c, nil, "Register success")
	}
}

// LoginHandler godoc
// @Security BearerAuth
// @Summary Login user
// @Tags Auth
// @Accept json
// @Produce json
// @Param data body LoginRequest true "Login data"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /login [post]
func (h *AuthHandler) LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.JSONError(c, 400, "Invalid request")
			return
		}
		if err := validate.Struct(req); err != nil {
			utils.JSONError(c, 400, err.Error())
			return
		}
		user, err := h.UserService.Authenticate(req.Email, req.Password)
		if err != nil {
			utils.JSONError(c, 401, "Invalid email or password")
			return
		}
		exp := time.Now().Add(24 * time.Hour)
		claims := &middleware.Claims{
			UserID: user.ID,
			Role:   user.Role,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(exp),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(middleware.JwtKey)
		if err != nil {
			utils.JSONError(c, 500, "Could not login")
			return
		}
		utils.JSONSuccess(c, gin.H{"token": tokenString}, "Login success")
	}
}

// MeHandler godoc
// @Security BearerAuth
// @Summary Get current user info
// @Tags Auth
// @Produce json
// @Success 200 {object} utils.SuccessResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /me [get]
func (h *AuthHandler) MeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("userID")
		user, err := h.UserService.FindByID(userID.(uint))
		if err != nil {
			utils.JSONError(c, 404, "User not found")
			return
		}
		utils.JSONSuccess(c, gin.H{"id": user.ID, "name": user.Name, "email": user.Email, "role": user.Role}, "User info")
	}
}
