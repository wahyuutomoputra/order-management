package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/wahyuutomoputra/order-management/models"
	"github.com/wahyuutomoputra/order-management/service"
	"github.com/wahyuutomoputra/order-management/utils"
)

var productValidate = validator.New()

type ProductRequest struct {
	Name  string  `json:"name" validate:"required,min=2"`
	Price float64 `json:"price" validate:"required,gt=0"`
	Stock int     `json:"stock" validate:"required,gte=0"`
}

type ProductHandler struct {
	ProductService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{ProductService: productService}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		if role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Admin only"})
			return
		}
		c.Next()
	}
}

// CreateProductHandler godoc
// @Summary Create product
// @Tags Product
// @Accept json
// @Produce json
// @Param data body ProductRequest true "Product data"
// @Success 201 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 403 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /admin/products [post]
// @Security BearerAuth
func (h *ProductHandler) CreateProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.JSONError(c, 400, "Invalid request")
			return
		}
		if err := productValidate.Struct(req); err != nil {
			utils.JSONError(c, 400, err.Error())
			return
		}
		product := models.Product{
			Name:  req.Name,
			Price: req.Price,
			Stock: req.Stock,
		}
		if err := h.ProductService.Create(&product); err != nil {
			utils.JSONError(c, 500, "Failed to create product")
			return
		}
		utils.JSONCreated(c, product, "Product created")
	}
}

// ListProductHandler godoc
// @Summary List products
// @Tags Product
// @Produce json
// @Success 200 {object} utils.SuccessResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /products [get]
func (h *ProductHandler) ListProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := h.ProductService.FindAll()
		if err != nil {
			utils.JSONError(c, 500, "Failed to get products")
			return
		}
		utils.JSONSuccess(c, products, "Product list")
	}
}

// GetProductHandler godoc
// @Summary Get product by ID
// @Tags Product
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /products/{id} [get]
func (h *ProductHandler) GetProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var id uint
		if err := parseUintParam(c, "id", &id); err != nil {
			utils.JSONError(c, 400, "Invalid product id")
			return
		}
		product, err := h.ProductService.FindByID(id)
		if err != nil {
			utils.JSONError(c, 404, "Product not found")
			return
		}
		utils.JSONSuccess(c, product, "Product detail")
	}
}

// UpdateProductHandler godoc
// @Summary Update product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param data body ProductRequest true "Product data"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 403 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /admin/products/{id} [put]
// @Security BearerAuth
func (h *ProductHandler) UpdateProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ProductRequest
		var id uint
		if err := parseUintParam(c, "id", &id); err != nil {
			utils.JSONError(c, 400, "Invalid product id")
			return
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.JSONError(c, 400, "Invalid request")
			return
		}
		if err := productValidate.Struct(req); err != nil {
			utils.JSONError(c, 400, err.Error())
			return
		}
		product, err := h.ProductService.FindByID(id)
		if err != nil {
			utils.JSONError(c, 404, "Product not found")
			return
		}
		product.Name = req.Name
		product.Price = req.Price
		product.Stock = req.Stock
		if err := h.ProductService.Update(product); err != nil {
			utils.JSONError(c, 500, "Failed to update product")
			return
		}
		utils.JSONSuccess(c, product, "Product updated")
	}
}

// DeleteProductHandler godoc
// @Summary Delete product
// @Tags Product
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 403 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /admin/products/{id} [delete]
// @Security BearerAuth
func (h *ProductHandler) DeleteProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var id uint
		if err := parseUintParam(c, "id", &id); err != nil {
			utils.JSONError(c, 400, "Invalid product id")
			return
		}
		product, err := h.ProductService.FindByID(id)
		if err != nil {
			utils.JSONError(c, 404, "Product not found")
			return
		}
		if err := h.ProductService.Delete(product); err != nil {
			utils.JSONError(c, 500, "Failed to delete product")
			return
		}
		utils.JSONSuccess(c, nil, "Product deleted")
	}
}

// parseUintParam helper
func parseUintParam(c *gin.Context, key string, out *uint) error {
	idStr := c.Param(key)
	var id64 uint64
	var err error
	if id64, err = strconv.ParseUint(idStr, 10, 32); err != nil {
		return err
	}
	*out = uint(id64)
	return nil
}
