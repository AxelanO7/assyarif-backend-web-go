package repository

import (
	"assyarif-backend-web-go/domain"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type posgreRtrRepository struct {
	DB *gorm.DB
}

func NewPostgreRtr(client *gorm.DB) domain.RtrRepository {
	return &posgreRtrRepository{
		DB: client,
	}
}

func (a *posgreRtrRepository) RetrieveRtrs() ([]domain.Rtr, error) {
	var res []domain.Rtr
	err := a.DB.
		Model(domain.Rtr{}).
		Preload("Outlet").
		Preload("Stock").
		Find(&res).Error
	if err != nil {
		return []domain.Rtr{}, err
	}
	fmt.Println(res)
	return res, nil
}

func (a *posgreRtrRepository) RetrieveRtrByID(id uint) (*domain.Rtr, error) {
	var res domain.Rtr
	err := a.DB.
		Model(domain.Rtr{}).
		Where("id = ?", id).
		Take(&res).Error
	if err != nil {
		return &domain.Rtr{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &domain.Rtr{}, fmt.Errorf("record not found")
	}
	fmt.Println(res)
	return &res, nil
}

func (a *posgreRtrRepository) CreateRtr(user *domain.Rtr) (*domain.Rtr, error) {
	err := a.DB.
		Model(domain.Rtr{}).
		Create(user).Error
	if err != nil {
		return &domain.Rtr{}, err
	}
	fmt.Println(user)
	return user, nil
}

func (a *posgreRtrRepository) UpdateRtr(user *domain.Rtr) (*domain.Rtr, error) {
	err := a.DB.
		Model(domain.Rtr{}).
		Where("id = ?", user.ID).
		Updates(user).Error
	if err != nil {
		return &domain.Rtr{}, err
	}
	fmt.Println(user)
	return user, nil
}

func (a *posgreRtrRepository) DeleteRtr(id uint) error {
	err := a.DB.
		Model(domain.Rtr{}).
		Where("id = ?", id).
		Delete(&domain.Rtr{}).Error
	if err != nil {
		return err
	}
	return nil
}
