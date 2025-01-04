package usecase

import (
	"assyarif-backend-web-go/domain"
	"context"
	"fmt"
	"time"
)

type stockUseCase struct {
	stockRepository domain.StockRepository
	contextTimeout  time.Duration
}

func NewStockUseCase(stock domain.StockRepository, t time.Duration) domain.StockUseCase {
	return &stockUseCase{
		stockRepository: stock,
		contextTimeout:  t,
	}
}

func (c *stockUseCase) FetchStockByID(ctx context.Context, id uint) (*domain.Stock, error) {
	res, err := c.stockRepository.RetrieveStockByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *stockUseCase) FetchStocks(ctx context.Context) ([]domain.Stock, error) {
	res, err := c.stockRepository.RetrieveAllStock()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *stockUseCase) CreateStock(ctx context.Context, req *domain.Stock) (*domain.Stock, error) {
	res, err := c.stockRepository.CreateStock(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *stockUseCase) UpdateStock(ctx context.Context, req *domain.Stock) (*domain.Stock, error) {
	res, err := c.stockRepository.UpdateStock(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *stockUseCase) DeleteStock(ctx context.Context, id uint) error {
	err := c.stockRepository.DeleteStock(id)
	if err != nil {
		return err
	}
	return nil
}

func (c *stockUseCase) DecreaseStocks(ctx context.Context, req []domain.Stock) ([]domain.Stock, error) {
	stocks, err := c.stockRepository.RetrieveAllStock()
	resultStocks := []domain.Stock{}
	if err != nil {
		return nil, err
	}
	for _, stock := range stocks {
		for _, reqStock := range req {
			if stock.IdStuff == reqStock.IdStuff {
				fmt.Println("match stock", stock, reqStock)
				reqStock.Quantity = stock.Quantity - reqStock.Quantity
				res, err := c.stockRepository.UpdateStockByStuffID(&reqStock)
				if err != nil {
					return nil, err
				}
				resultStocks = append(resultStocks, *res)
			}
		}
	}
	afterStocks, err := c.stockRepository.RetrieveAllStock()
	if err != nil {
		return nil, err
	}
	fmt.Println("afterStocks", afterStocks)
	return resultStocks, nil
}

func (c *stockUseCase) IncreaseStocks(ctx context.Context, req []domain.Stock) ([]domain.Stock, error) {
	stocks, err := c.stockRepository.RetrieveAllStock()
	resultStocks := []domain.Stock{}
	if err != nil {
		return nil, err
	}
	for _, stock := range stocks {
		for _, reqStock := range req {
			if stock.IdStuff == reqStock.IdStuff {
				reqStock.Quantity = stock.Quantity + reqStock.Quantity
				res, err := c.stockRepository.UpdateStockByStuffID(&reqStock)
				if err != nil {
					return nil, err
				}
				resultStocks = append(resultStocks, *res)
			}
		}
	}
	afterStocks, err := c.stockRepository.RetrieveAllStock()
	if err != nil {
		return nil, err
	}
	fmt.Println("afterStocks", afterStocks)
	return resultStocks, nil
}

func (c *stockUseCase) UpdateDescription(ctx context.Context, req []domain.UpdateDescriptionRequest) ([]domain.Stock, error) {
	stocks, err := c.stockRepository.RetrieveAllStock()
	resultStocks := []domain.Stock{}
	if err != nil {
		return nil, err
	}
	for _, stock := range stocks {
		for _, reqStock := range req {
			if stock.IdStuff == reqStock.ID {
				fmt.Println("match stock", stock, reqStock)
				stock.Description = &reqStock.Description
				res, err := c.stockRepository.UpdateDescription(&reqStock)
				if err != nil {
					return nil, err
				}
				resultStocks = append(resultStocks, *res)
			}
		}
	}
	return resultStocks, nil
}
