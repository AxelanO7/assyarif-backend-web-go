package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Rtr struct {
	ID          uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	OutletID    uint           `json:"outlet_id"`
	Outlet      Outlet         `gorm:"foreignKey:OutletID" json:"outlet"`
	StockID     uint           `json:"stock_id"`
	Stock       Stock          `gorm:"foreignKey:StockID" json:"stock"`
	TotalReturn float64        `json:"total_return"`
	Reason      string         `json:"reason"`
	Proof       string         `json:"proof"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type PeriodRtr struct {
	Date string `json:"date"`
	Rtrs []Rtr  `json:"rtrs"`
}

type RtrRepository interface {
	RetrieveRtrs() ([]Rtr, error)
	RetrieveRtrByID(id uint) (*Rtr, error)
	CreateRtr(Rtr *Rtr) (*Rtr, error)
	UpdateRtr(Rtr *Rtr) (*Rtr, error)
	DeleteRtr(id uint) error
}

type RtrUseCase interface {
	FetchRtrs(ctx context.Context) ([]Rtr, error)
	FetchRtrByID(ctx context.Context, id uint) (*Rtr, error)
	CreateRtr(ctx context.Context, req *Rtr) (*Rtr, error)
	UpdateRtr(ctx context.Context, req *Rtr) (*Rtr, error)
	DeleteRtr(ctx context.Context, id uint) error
	GetRtrsByPeriod(ctx context.Context) ([]PeriodRtr, error)
}
