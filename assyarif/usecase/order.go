package usecase

import (
	"assyarif-backend-web-go/domain"
	"context"
	"fmt"
	"time"
)

type orderUseCase struct {
	orderRepository domain.OrderRepository
	outRepository   domain.OutRepository
	contextTimeout  time.Duration
}

func NewOrderUseCase(order domain.OrderRepository, out domain.OutRepository,
	t time.Duration) domain.OrderUseCase {
	return &orderUseCase{
		orderRepository: order,
		outRepository:   out,
		contextTimeout:  t,
	}
}

func (c *orderUseCase) ShowOrders(ctx context.Context) ([]domain.Order, error) {
	// return c.orderRepository.RetrieveOrders()
	finalOrders := []domain.Order{}
	resOrders, err := c.orderRepository.RetrieveOrders()
	if err != nil {
		return nil, err
	}
	// get all outs
	finalOrders = resOrders
	outs, err := c.outRepository.RetrieveOuts()
	if err != nil {
		return finalOrders, nil
	}
	// for _, order := range resOrders {
	// 	for _, out := range outs {
	// 		if order.ID != out.Order.ID {
	// 			finalOrders = append(finalOrders, order)
	// 		}
	// 	}
	// }
	// // remove duplicate
	// for i := 0; i < len(finalOrders); i++ {
	// 	for j := i + 1; j < len(finalOrders); j++ {
	// 		if finalOrders[i].ID == finalOrders[j].ID {
	// 			finalOrders = append(finalOrders[:j], finalOrders[j+1:]...)
	// 			j--
	// 		}
	// 	}
	// }
	for i := 0; i < len(finalOrders); i++ {
		for j := 0; j < len(outs); j++ {
			if finalOrders[i].ID == outs[j].Order.ID {
				finalOrders = append(finalOrders[:i], finalOrders[i+1:]...)
				i--
			}
		}
	}
	return finalOrders, nil
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
	return c.orderRepository.DeleteOrdersById(in)
}
