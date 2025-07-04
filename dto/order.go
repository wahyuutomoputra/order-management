package dto

// OrderItemInput adalah DTO untuk input item order
// Digunakan pada proses pembuatan order

type OrderItemInput struct {
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  int  `json:"quantity" validate:"required,gt=0"`
}

// OrderRequest adalah DTO untuk request pembuatan order
// Berisi daftar item yang ingin diorder

type OrderRequest struct {
	Items []OrderItemInput `json:"items" validate:"required,dive"`
}
