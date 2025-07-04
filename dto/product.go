package dto

// ProductRequest adalah DTO untuk request pembuatan/ubah produk

type ProductRequest struct {
	Name  string  `json:"name" validate:"required,min=2"`
	Price float64 `json:"price" validate:"required,gt=0"`
	Stock int     `json:"stock" validate:"required,gte=0"`
}
