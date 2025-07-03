package utils

import "github.com/gin-gonic/gin"

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func JSONSuccess(c *gin.Context, data interface{}, message string) {
	c.JSON(200, SuccessResponse{Success: true, Data: data, Message: message})
}

func JSONCreated(c *gin.Context, data interface{}, message string) {
	c.JSON(201, SuccessResponse{Success: true, Data: data, Message: message})
}

func JSONError(c *gin.Context, code int, err string) {
	c.JSON(code, ErrorResponse{Success: false, Error: err})
}
