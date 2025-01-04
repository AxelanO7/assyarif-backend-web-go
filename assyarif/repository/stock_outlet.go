package repository

import (
	"assyarif-backend-web-go/domain"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type posgreStockOutletRepository struct {
	DB *gorm.DB
}

func NewPostgreStockOutlet(client *gorm.DB) domain.StockOutletRepository {
	return &posgreStockOutletRepository{
		DB: client,
	}
}

func (a *posgreStockOutletRepository) RetrieveAllStockOutlet() ([]domain.StockOutlet, error) {
	var res []domain.StockOutlet
	err := a.DB.
		Model(domain.StockOutlet{}).
		Preload("Out").
		Preload("Outlet").
		Find(&res).Error
	if err != nil {
		return []domain.StockOutlet{}, err
	}
	fmt.Println(res)
	return res, nil
}

func (a *posgreStockOutletRepository) RetrieveStockOutletByID(id uint) (*domain.StockOutlet, error) {
	var res domain.StockOutlet
	err := a.DB.
		Model(domain.StockOutlet{}).
		Where("id = ?", id).
		Take(&res).Error
	if err != nil {
		return &domain.StockOutlet{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.StockOutlet{}, fmt.Errorf("record not found")
	}
	fmt.Println(res)
	return &res, nil
}

func (a *posgreStockOutletRepository) CreateStockOutlet(user *domain.StockOutlet) (*domain.StockOutlet, error) {
	err := a.DB.
		Model(domain.StockOutlet{}).
		Create(user).Error
	if err != nil {
		return &domain.StockOutlet{}, err
	}
	fmt.Println(user)
	return user, nil
}

func (a *posgreStockOutletRepository) UpdateStockOutlet(user *domain.StockOutlet) (*domain.StockOutlet, error) {
	err := a.DB.
		Model(domain.StockOutlet{}).
		Where("id = ?", user.ID).
		Updates(user).Error
	if err != nil {
		return &domain.StockOutlet{}, err
	}
	fmt.Println(user)
	return user, nil
}

func (a *posgreStockOutletRepository) DeleteStockOutlet(id uint) error {
	err := a.DB.
		Model(domain.StockOutlet{}).
		Where("id = ?", id).
		Delete(&domain.StockOutlet{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *posgreStockOutletRepository) UpdateStockOutletsMultiple(user []domain.StockOutlet) ([]domain.StockOutlet, error) {
	fmt.Println("masuk repository update")
	for _, u := range user {
		err := a.DB.
			Model(domain.StockOutlet{}).
			Where("id = ?", u.ID).
			Updates(u).Error
		if err != nil {
			return []domain.StockOutlet{}, err
		}
	}
	fmt.Println(user)
	return user, nil
}

func (a *posgreStockOutletRepository) CreateStockOutletsMultiple(user []domain.StockOutlet) ([]domain.StockOutlet, error) {
	fmt.Println("masuk repository create")
	for _, u := range user {
		err := a.DB.
			Model(domain.StockOutlet{}).
			Create(u).Error
		if err != nil {
			return []domain.StockOutlet{}, err
		}
	}
	fmt.Println(user)
	return user, nil
}

// func (a *posgreStockOutletRepository) IncreaseDashboard(user *domain.StockOutlet) (*domain.StockOutlet, error) {
// 	var res domain.StockOutlet
// 	err := a.DB.
// 		Model(domain.StockOutlet{}).
// 		Where("id_stuff = ?", user.IdStuff).
// 		Take(&res).Error
// 	if err != nil {
// 		return &domain.StockOutlet{}, err
// 	}
// 	if errors.Is(err, gorm.ErrRecordNotFound) {
// 		return &domain.StockOutlet{}, fmt.Errorf("record not found")
// 	}
// 	user.Quantity = user.Quantity + res.Quantity
// 	err = a.DB.
// 		Model(domain.StockOutlet{}).
// 		Where("id = ?", res.ID).
// 		Updates(user).Error
// 	if err != nil {
// 		return &domain.StockOutlet{}, err
// 	}
// 	fmt.Println(user)
// 	return user, nil
// }

// func (a *posgreStockOutletRepository) DecreaseDashboard(user *domain.StockOutlet) (*domain.StockOutlet, error) {
// 	var res domain.StockOutlet
// 	err := a.DB.
// 		Model(domain.StockOutlet{}).
// 		Where("id_stuff = ?", user.IdStuff).
// 		Take(&res).Error
// 	if err != nil {
// 		return &domain.StockOutlet{}, err
// 	}
// 	if errors.Is(err, gorm.ErrRecordNotFound) {
// 		return &domain.StockOutlet{}, fmt.Errorf("record not found")
// 	}
// 	user.Quantity = res.Quantity - user.Quantity
// 	err = a.DB.
// 		Model(domain.StockOutlet{}).
// 		Where("id = ?", res.ID).
// 		Updates(user).Error
// 	if err != nil {
// 		return &domain.StockOutlet{}, err
// 	}
// 	fmt.Println(user)
// 	return user, nil
// }
