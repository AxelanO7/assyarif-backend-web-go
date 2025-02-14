package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Out struct {
	ID          uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	OutID       uint           `json:"out_id"`
	OrderID     uint           `json:"order_id"`
	Order       Order          `gorm:"foreignKey:OrderID" json:"order"`
	TotalPaided float64        `json:"total_paided" default:"0"`
	ReturnCash  float64        `json:"return_cash" default:"0"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type PeriodOut struct {
	Date string `json:"date"`
	Outs []Out  `json:"outs"`
}

type OutRepository interface {
	RetrieveOuts() ([]Out, error)
	RetrieveOutByID(id string) (Out, error)
	CreateOut(out Out) (Out, error)
	UpdateOutByID(out Out) (Out, error)
	RemoveOutByID(id string) error
	CreateOuts(outs []Out) ([]Out, error)
}

type OutUseCase interface {
	GetOuts(ctx context.Context) ([]Out, error)
	ShowOutByID(ctx context.Context, id string) (Out, error)
	AddOut(ctx context.Context, out Out) (Out, error)
	EditOutByID(ctx context.Context, out Out) (Out, error)
	DeleteOutByID(ctx context.Context, id string) error
	ShowOutLastOrderID(ctx context.Context) (int, error)
	AddOuts(ctx context.Context, outs []Out) ([]Out, error)
	GetOutsByPeriod(ctx context.Context) ([]PeriodOut, error)
}
