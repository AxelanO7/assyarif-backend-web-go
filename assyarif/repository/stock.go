package repository

import (
	"assyarif-backend-web-go/domain"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type posgreStockRepository struct {
	DB *gorm.DB
}

func NewPostgreStock(client *gorm.DB) domain.StockRepository {
	return &posgreStockRepository{
		DB: client,
	}
}

func (a *posgreStockRepository) RetrieveStocks() ([]domain.Stock, error) {
	var res []domain.Stock
	err := a.DB.
		Model(domain.Stock{}).
		Find(&res).Error
	if err != nil {
		return []domain.Stock{}, err
	}
	fmt.Println(res)
	return res, nil
}

func (a *posgreStockRepository) RetrieveStockByID(id uint) (*domain.Stock, error) {
	var res domain.Stock
	err := a.DB.
		Model(domain.Stock{}).
		Where("id = ?", id).
		Take(&res).Error
	if err != nil {
		return &domain.Stock{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.Stock{}, fmt.Errorf("record not found")
	}
	fmt.Println(res)
	return &res, nil
}

func (a *posgreStockRepository) CreateStock(user *domain.Stock) (*domain.Stock, error) {
	err := a.DB.
		Model(domain.Stock{}).
		Create(user).Error
	if err != nil {
		return &domain.Stock{}, err
	}
	fmt.Println(user)
	return user, nil
}

func (a *posgreStockRepository) UpdateStock(stock *domain.Stock) (*domain.Stock, error) {
	err := a.DB.
		Model(domain.Stock{}).
		Where("id = ?", stock.ID).
		Updates(stock).Error
	if err != nil {
		return &domain.Stock{}, err
	}
	fmt.Println("updated stock", stock)
	return stock, nil
}

func (a *posgreStockRepository) DeleteStock(id uint) error {
	err := a.DB.
		Model(domain.Stock{}).
		Where("id = ?", id).
		Delete(&domain.Stock{}).Error
	if err != nil {
		return err

	}
	return nil
}

func (a *posgreStockRepository) UpdateStockByStuffID(stock *domain.Stock) (*domain.Stock, error) {
	err := a.DB.
		Model(domain.Stock{}).
		Where("id_stuff = ?", stock.IdStuff).
		Update("quantity", stock.Quantity).Error
	if err != nil {
		return &domain.Stock{}, err
	}
	fmt.Println("updated stock", stock)
	return stock, nil
}

func (a *posgreStockRepository) UpdateDescription(req *domain.UpdateDescriptionRequest) (*domain.Stock, error) {
	fmt.Println("req", req)
	err := a.DB.
		Model(domain.Stock{}).
		Where("id_stuff = ?", req.ID).
		Update("description", req.Description).Error
	if err != nil {
		return &domain.Stock{}, err
	}
	fmt.Println("updated stock", req)
	return &domain.Stock{}, nil
}
