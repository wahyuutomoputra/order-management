package models

type Order struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	CreatedAt int64
	Items     []OrderItem
}

type OrderItem struct {
	ID        uint `gorm:"primaryKey"`
	OrderID   uint
	ProductID uint
	Quantity  int
	Price     float64 // harga saat order
}
