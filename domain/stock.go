package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Stock struct {
	ID          uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	IdStuff     uint           `json:"id_stuff"`
	Name        string         `json:"name"`
	Type        string         `json:"type"`
	Quantity    float64        `json:"quantity"`
	Unit        string         `json:"unit"`
	Price       float64        `json:"price"`
	Description *string        `json:"description"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type UpdateDescriptionRequest struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
}

type PeriodStock struct {
	Date   string        `json:"date"`
	Stocks []StockReport `json:"stocks"`
}
type StockReport struct {
	ID           uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	IdStuff      uint           `json:"id_stuff"`
	Name         string         `json:"name"`
	Type         string         `json:"type"`
	Quantity     float64        `json:"quantity"`
	Unit         string         `json:"unit"`
	Price        float64        `json:"price"`
	Description  *string        `json:"description"`
	InitialStock float64        `json:"initial_stock"`
	FinalStock   float64        `json:"final_stock"`
	InStock      float64        `json:"in_stock"`
	OutStock     float64        `json:"out_stock"`
	CreatedAt    *time.Time     `json:"created_at"`
	UpdatedAt    *time.Time     `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type StockRepository interface {
	RetrieveStocks() ([]Stock, error)
	RetrieveStockByID(id uint) (*Stock, error)
	CreateStock(Stock *Stock) (*Stock, error)
	UpdateStock(Stock *Stock) (*Stock, error)
	UpdateStockByStuffID(Stok *Stock) (*Stock, error)
	DeleteStock(id uint) error
	UpdateDescription(req *UpdateDescriptionRequest) (*Stock, error)
}

type StockUseCase interface {
	FetchStocks(ctx context.Context) ([]Stock, error)
	FetchStockByID(ctx context.Context, id uint) (*Stock, error)
	CreateStock(ctx context.Context, req *Stock) (*Stock, error)
	UpdateStock(ctx context.Context, req *Stock) (*Stock, error)
	DeleteStock(ctx context.Context, id uint) error
	IncreaseStocks(ctx context.Context, req []Stock) ([]Stock, error)
	DecreaseStocks(ctx context.Context, req []Stock) ([]Stock, error)
	UpdateDescription(ctx context.Context, req []UpdateDescriptionRequest) ([]Stock, error)
	GetStocksByPeriod(ctx context.Context) ([]PeriodStock, error)
}
