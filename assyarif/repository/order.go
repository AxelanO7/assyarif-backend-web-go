package repository

import (
	"assyarif-backend-web-go/domain"
	"fmt"

	"gorm.io/gorm"
)

type posgreOrderRepository struct {
	DB *gorm.DB
}

func NewPostgreOrder(client *gorm.DB) domain.OrderRepository {
	return &posgreOrderRepository{
		DB: client,
	}
}

func (a *posgreOrderRepository) RetrieveOrders() ([]domain.Order, error) {
	var res []domain.Order
	err := a.DB.
		Model(domain.Order{}).
		Preload("Outlet").
		Preload("Stock").
		Find(&res).Error
	if err != nil {
		return []domain.Order{}, err
	}
	fmt.Println(res)
	return res, nil
}

func (a *posgreOrderRepository) CreateOrder(order domain.Order) (domain.Order, error) {
	err := a.DB.
		Model(domain.Order{}).
		Create(&order).Error
	if err != nil {
		return domain.Order{}, err
	}
	fmt.Println(order)
	return order, nil
}

func (a *posgreOrderRepository) RetrieveOrderById(id string) (domain.Order, error) {
	var res domain.Order
	err := a.DB.
		Model(domain.Order{}).
		Where("id = ?", id).
		Preload("Outlet").
		Preload("Stock").
		First(&res).Error
	if err != nil {
		return domain.Order{}, err
	}
	fmt.Println(res)
	return res, nil
}

func (a *posgreOrderRepository) UpdateOrderById(order domain.Order) (domain.Order, error) {
	err := a.DB.
		Model(domain.Order{}).
		Where("id = ?", order.ID).
		Updates(&order).Error
	if err != nil {
		return domain.Order{}, err
	}
	fmt.Println(order)
	return order, nil
}

func (a *posgreOrderRepository) RemoveOrderById(id string) error {
	err := a.DB.
		Model(domain.Order{}).
		Where("id = ?", id).
		Delete(&domain.Order{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *posgreOrderRepository) RetrieveOrderByOutletId(id string) ([]domain.Order, error) {
	var res []domain.Order
	err := a.DB.
		Model(domain.Order{}).
		Where("outlet_id = ?", id).
		Preload("Outlet").
		Preload("Stock").
		Find(&res).Error
	if err != nil {
		return []domain.Order{}, err
	}
	if len(res) == 0 {
		return []domain.Order{}, nil
	}
	fmt.Println(res)
	return res, nil
}

func (a *posgreOrderRepository) CreateOrders(order []domain.Order) ([]domain.Order, error) {
	err := a.DB.
		Model(domain.Order{}).
		Create(&order).Error
	if err != nil {
		return []domain.Order{}, err
	}
	fmt.Println(order)
	return order, nil
}

func (a *posgreOrderRepository) DeleteOrdersById(in []uint) error {
	fmt.Println("DeleteOrdersById Repository", in)
	err := a.DB.
		Model(domain.Order{}).
		Where("id IN ?", in).
		Delete(&domain.Order{}).Error
	if err != nil {
		return err
	}
	return nil
}
