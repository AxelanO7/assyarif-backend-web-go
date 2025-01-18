package usecase

import (
	"assyarif-backend-web-go/domain"
	"context"
	"fmt"
	"time"
)

type orderUseCase struct {
	orderRepository domain.OrderRepository
	contextTimeout  time.Duration
}

func NewOrderUseCase(order domain.OrderRepository, t time.Duration) domain.OrderUseCase {
	return &orderUseCase{
		orderRepository: order,
		contextTimeout:  t,
	}
}

func (c *orderUseCase) ShowOrders(ctx context.Context) ([]domain.Order, error) {
	return c.orderRepository.RetrieveOrders()
}

func (c *orderUseCase) AddOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	return c.orderRepository.CreateOrder(order)
}

func (c *orderUseCase) ShowOrderById(ctx context.Context, id string) (domain.Order, error) {
	return c.orderRepository.RetrieveOrderById(id)
}

func (c *orderUseCase) EditOrderById(ctx context.Context, order domain.Order) (domain.Order, error) {
	return c.orderRepository.UpdateOrderById(order)
}

func (c *orderUseCase) DeleteOrderById(ctx context.Context, id string) error {
	return c.orderRepository.RemoveOrderById(id)
}

func (c *orderUseCase) ShowOrderByOutletId(ctx context.Context, id string) ([]domain.Order, error) {
	return c.orderRepository.RetrieveOrderByOutletId(id)
}

func (c *orderUseCase) AddOrders(ctx context.Context, orders []domain.Order) ([]domain.Order, error) {
	return c.orderRepository.CreateOrders(orders)
}

func (c *orderUseCase) DeleteOrdersById(ctx context.Context, in []uint) error {
	fmt.Println("DeleteOrdersById UseCase", in)
	return c.orderRepository.DeleteOrdersById(in)
}
