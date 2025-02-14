package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type In struct {
	ID        uint           `gorm:"primarykey;AUTO_INCREMENT" json:"id"`
	IdStuff   uint           `json:"id_stuff"`
	Name      string         `json:"name"`
	Type      string         `json:"type"`
	Quantity  float64        `json:"quantity"`
	Unit      string         `json:"unit"`
	Price     float64        `json:"price"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type PeriodIn struct {
	Date string `json:"date"`
	Ins  []In   `json:"ins"`
}

type InRepository interface {
	RetrieveIns() ([]In, error)
	RetrieveInById(id string) (In, error)
	CreateIn(in In) (In, error)
	UpdateInById(in In) (In, error)
	RemoveInById(id string) error
}

type InUseCase interface {
	ShowIns(ctx context.Context) ([]In, error)
	ShowInById(ctx context.Context, id string) (In, error)
	ShowInLastNumber(ctx context.Context) (int, error)
	AddIn(ctx context.Context, in In) (In, error)
	EditInById(ctx context.Context, in In) (In, error)
	DeleteInById(ctx context.Context, id string) error
	GetInsByPeriod(ctx context.Context) ([]PeriodIn, error)
}
