package usecase

import (
	"assyarif-backend-web-go/domain"
	"context"
	"fmt"
	"time"
)

type stockOutletUseCase struct {
	stockOutletRepository domain.StockOutletRepository
	contextTimeout        time.Duration
}

func NewStockOutletUseCase(stockOutlet domain.StockOutletRepository, t time.Duration) domain.StockOutletUseCase {
	return &stockOutletUseCase{
		stockOutletRepository: stockOutlet,
		contextTimeout:        t,
	}
}

func (c *stockOutletUseCase) FetchStockOutletByID(ctx context.Context, id uint) (*domain.StockOutlet, error) {
	res, err := c.stockOutletRepository.RetrieveStockOutletByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *stockOutletUseCase) FetchStockOutlets(ctx context.Context) ([]domain.StockOutlet, error) {
	res, err := c.stockOutletRepository.RetrieveAllStockOutlet()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *stockOutletUseCase) CreateStockOutlet(ctx context.Context, req *domain.StockOutlet) (*domain.StockOutlet, error) {
	res, err := c.stockOutletRepository.CreateStockOutlet(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *stockOutletUseCase) UpdateStockOutlet(ctx context.Context, req *domain.StockOutlet) (*domain.StockOutlet, error) {
	res, err := c.stockOutletRepository.UpdateStockOutlet(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *stockOutletUseCase) DeleteStockOutlet(ctx context.Context, id uint) error {
	err := c.stockOutletRepository.DeleteStockOutlet(id)
	if err != nil {
		return err
	}
	return nil
}

func (c *stockOutletUseCase) IncreaseDashboard(ctx context.Context, req *domain.StockOutlet) (*domain.StockOutlet, error) {
	stocks, errStocks := c.stockOutletRepository.RetrieveAllStockOutlet()
	if errStocks != nil {
		return nil, errStocks
	}
	for _, stock := range stocks {
		if stock.IdStuff == req.IdStuff {
			req = &domain.StockOutlet{
				ID:        stock.ID,
				IdStuff:   stock.IdStuff,
				IdOut:     stock.IdOut,
				Out:       stock.Out,
				Name:      stock.Name,
				Type:      stock.Type,
				Quantity:  stock.Quantity + req.Quantity,
				Unit:      stock.Unit,
				Price:     stock.Price,
				IdOutlet:  stock.IdOutlet,
				CreatedAt: stock.CreatedAt,
				UpdatedAt: stock.UpdatedAt,
				DeletedAt: stock.DeletedAt,
			}
			res, err := c.stockOutletRepository.UpdateStockOutlet(req)
			if err != nil {
				return nil, err
			}
			return res, nil
		}
	}
	res, err := c.stockOutletRepository.CreateStockOutlet(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *stockOutletUseCase) DecreaseDashboard(ctx context.Context, req *domain.StockOutlet) (*domain.StockOutlet, error) {
	stocks, errStocks := c.stockOutletRepository.RetrieveAllStockOutlet()
	if errStocks != nil {
		return nil, errStocks
	}
	for _, stock := range stocks {
		if stock.IdStuff == req.IdStuff {
			req = &domain.StockOutlet{
				ID:        stock.ID,
				IdStuff:   stock.IdStuff,
				IdOut:     stock.IdOut,
				Out:       stock.Out,
				Name:      stock.Name,
				Type:      stock.Type,
				Quantity:  stock.Quantity - req.Quantity,
				Unit:      stock.Unit,
				Price:     stock.Price,
				IdOutlet:  stock.IdOutlet,
				CreatedAt: stock.CreatedAt,
				UpdatedAt: stock.UpdatedAt,
				DeletedAt: stock.DeletedAt,
			}
			res, err := c.stockOutletRepository.UpdateStockOutlet(req)
			if err != nil {
				return nil, err
			}
			return res, nil
		}
	}
	return nil, nil
}

func (c *stockOutletUseCase) IncreaseDashboardMultiple(ctx context.Context, req []domain.StockOutlet) ([]domain.StockOutlet, error) {
	stockOutlets := []domain.StockOutlet{}
	fmt.Println("masuk usecase")
	for _, stockReq := range req {
		fmt.Println("masuk usecase loop")
		// stocks := c.stockOutletRepository.RetrieveAllStockOutlet()
		// if stocks == nil {
		// 	fmt.Println("StockReq1", stockReq)
		// 	st, err := c.stockOutletRepository.CreateStockOutlet(&stockReq)
		// 	if err != nil {
		// 		fmt.Println("StockReq2", stockReq, "Stock2", st)
		// 		return nil, err
		// 	}
		// 	stockOutlets = append(stockOutlets, *st)
		// }
		// if stock length 0
		// if errStocks != nil {
		// 	fmt.Println("StockReq1", stockReq)
		// 	st, err := c.stockOutletRepository.CreateStockOutlet(&stockReq)
		// 	if err != nil {
		// 		fmt.Println("StockReq2", stockReq, "Stock2", st)
		// 		return nil, err
		// 	}
		// 	stockOutlets = append(stockOutlets, *st)
		// }

		stocks, errStocks := c.stockOutletRepository.RetrieveAllStockOutlet()
		fmt.Println("StockReq1", stockReq)
		fmt.Println("Stocks", stocks)
		if errStocks != nil {
			fmt.Println("StockReq1", stockReq)
			st, err := c.stockOutletRepository.CreateStockOutlet(&stockReq)
			if err != nil {
				fmt.Println("StockReq2", stockReq, "Stock2", st)
				return nil, err
			}
			stockOutlets = append(stockOutlets, *st)
		}
		if len(stocks) == 0 {
			fmt.Println("StockReq1", stockReq)
			st, err := c.stockOutletRepository.CreateStockOutlet(&stockReq)
			if err != nil {
				fmt.Println("StockReq2", stockReq, "Stock2", st)
				return nil, err
			}
			stockOutlets = append(stockOutlets, *st)
		}
		for _, st := range stocks {
			fmt.Println("StockReq", stockReq, "Stock", st)
			if st.IdStuff == stockReq.IdStuff {
				fmt.Println("StockReq Match", stockReq, "Stock Match", st)
				increaseStock := st.Quantity + stockReq.Quantity
				stockReq = domain.StockOutlet{
					ID:        st.ID,
					IdStuff:   st.IdStuff,
					IdOut:     stockReq.IdOut,
					Out:       st.Out,
					Name:      st.Name,
					Type:      st.Type,
					Quantity:  increaseStock,
					Unit:      st.Unit,
					Price:     st.Price,
					IdOutlet:  st.IdOutlet,
					CreatedAt: st.CreatedAt,
					UpdatedAt: st.UpdatedAt,
					DeletedAt: st.DeletedAt,
				}
				res, err := c.stockOutletRepository.UpdateStockOutlet(&stockReq)
				if err != nil {
					return nil, err
				}
				_, errCreated := c.stockOutletRepository.CreateStockOutlet(&stockReq)
				if errCreated != nil {
					return nil, errCreated
				}
				stockOutlets = append(stockOutlets, *res)
			}
		}
	}
	return stockOutlets, nil
}
