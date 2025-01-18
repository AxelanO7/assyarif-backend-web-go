package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID         uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	OutletID   uint           `json:"outlet_id"`
	Outlet     Outlet         `gorm:"foreignKey:OutletID" json:"outlet"`
	StockID    uint           `json:"stock_id"`
	Stock      Stock          `gorm:"foreignKey:StockID" json:"stock"`
	TotalPaid  float64        `json:"total_paid"`
	TotalOrder float64        `json:"total_order"`
	Status     int32          `json:"status"`
	CreatedAt  *time.Time     `json:"created_at"`
	UpdatedAt  *time.Time     `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type OrderRepository interface {
	RetrieveOrders() ([]Order, error)
	RetrieveOrderById(id string) (Order, error)
	CreateOrder(in Order) (Order, error)
	UpdateOrderById(in Order) (Order, error)
	RemoveOrderById(id string) error
	RetrieveOrderByOutletId(id string) ([]Order, error)
	CreateOrders(in []Order) ([]Order, error)
	DeleteOrdersById(in []uint) error
}

type OrderUseCase interface {
	ShowOrders(ctx context.Context) ([]Order, error)
	ShowOrderById(ctx context.Context, id string) (Order, error)
	AddOrder(ctx context.Context, in Order) (Order, error)
	EditOrderById(ctx context.Context, in Order) (Order, error)
	DeleteOrderById(ctx context.Context, id string) error
	ShowOrderByOutletId(ctx context.Context, id string) ([]Order, error)
	AddOrders(ctx context.Context, in []Order) ([]Order, error)
	DeleteOrdersById(ctx context.Context, in []uint) error
}
