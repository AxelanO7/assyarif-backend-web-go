package usecase

import (
	"assyarif-backend-web-go/domain"
	"assyarif-backend-web-go/middleware"
	"assyarif-backend-web-go/utils"
	"context"
	"fmt"
	"time"
)

type userUseCase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUseCase(user domain.UserRepository, t time.Duration) domain.UserUseCase {
	return &userUseCase{
		userRepository: user,
		contextTimeout: t,
	}
}

func (c *userUseCase) LoginUser(ctx context.Context, req *domain.LoginPayload) (*domain.User, string, error) {
	res, err := c.userRepository.RetrieveByUsername(*req.Username)
	if err != nil {
		return nil, "", err
	}
	err = utils.VerifyPassword(req.Password, res.Password)
	if err != nil {
		return nil, "", fmt.Errorf("error verifying password: %v", err)
	}
	tokPay := domain.TokenClaims{
		User: res,
	}
	token, err := middleware.CreateToken(&tokPay)
	if err != nil {
		return nil, "", fmt.Errorf("cannot create token: %v", err)
	}
	return res, token, nil
}

func (c *userUseCase) FetchUserByID(ctx context.Context, id uint) (*domain.User, error) {
	res, err := c.userRepository.RetrieveUserByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *userUseCase) FetchUsers(ctx context.Context) ([]domain.User, error) {
	res, err := c.userRepository.RetrieveAllUser()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *userUseCase) CreateUser(ctx context.Context, req *domain.User) (*domain.User, error) {
	password, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("unable to hash password: %v", err)
	}

	req.Password = password
	res, err := c.userRepository.CreateUser(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *userUseCase) UpdateUser(ctx context.Context, req *domain.User) (*domain.User, error) {
	res, err := c.userRepository.UpdateUser(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *userUseCase) DeleteUser(ctx context.Context, id uint) error {
	err := c.userRepository.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (c *userUseCase) ShowUserLastNumber(ctx context.Context) (int, error) {
	var res []domain.User
	res, err := c.userRepository.RetrieveAllUser()
	if err != nil {
		return 0, err
	}

	lastNumber := 0
	for _, v := range res {
		fmt.Println(v.ID)
		if v.ID > uint(lastNumber) {
			lastNumber = int(v.ID)
		}

	}

	fmt.Println(lastNumber)
	return lastNumber, nil
}
