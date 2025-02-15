package usecase

import (
	"assyarif-backend-web-go/domain"
	"context"
	"fmt"
	"time"
)

type stockUseCase struct {
	stockRepository  domain.StockRepository
	insRepository    domain.InRepository
	outsRepository   domain.OutRepository
	returnRepository domain.RtrRepository
	contextTimeout   time.Duration
}

func NewStockUseCase(
	stock domain.StockRepository,
	in domain.InRepository,
	outs domain.OutRepository,
	returns domain.RtrRepository,
	t time.Duration) domain.StockUseCase {
	return &stockUseCase{
		stockRepository:  stock,
		insRepository:    in,
		outsRepository:   outs,
		returnRepository: returns,
		contextTimeout:   t,
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
	res, err := c.stockRepository.RetrieveStocks()
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
	stocks, err := c.stockRepository.RetrieveStocks()
	resultStocks := []domain.Stock{}
	if err != nil {
		return nil, err
	}
	for _, stock := range stocks {
		for _, reqStock := range req {
			if stock.IdStuff == reqStock.IdStuff {
				reqStock.Quantity = stock.Quantity - reqStock.Quantity
				res, err := c.stockRepository.UpdateStockByStuffID(&reqStock)
				if err != nil {
					return nil, err
				}
				resultStocks = append(resultStocks, *res)
			}
		}
	}
	afterStocks, err := c.stockRepository.RetrieveStocks()
	if err != nil {
		return nil, err
	}
	fmt.Println("afterStocks", afterStocks)
	return resultStocks, nil
}

func (c *stockUseCase) IncreaseStocks(ctx context.Context, req []domain.Stock) ([]domain.Stock, error) {
	stocks, err := c.stockRepository.RetrieveStocks()
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
	afterStocks, err := c.stockRepository.RetrieveStocks()
	if err != nil {
		return nil, err
	}
	fmt.Println("afterStocks", afterStocks)
	return resultStocks, nil
}

func (c *stockUseCase) UpdateDescription(ctx context.Context, req []domain.UpdateDescriptionRequest) ([]domain.Stock, error) {
	stocks, err := c.stockRepository.RetrieveStocks()
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

func (c *stockUseCase) GetStocksByPeriod(ctx context.Context) ([]domain.PeriodStock, error) {
	// get stocks
	stocks, err := c.stockRepository.RetrieveStocks()
	if err != nil {
		return nil, err
	}

	// assign stock to period
	finalPeriodStock := []domain.StockReport{}
	stocksByPeriod := make(map[string][]domain.Stock)
	for _, stock := range stocks {
		period := fmt.Sprintf("%d-%d", stock.CreatedAt.Month(), stock.CreatedAt.Year())
		stocksByPeriod[period] = append(stocksByPeriod[period], stock)
		// add to final stock
		finalPeriodStock = append(finalPeriodStock, domain.StockReport{
			ID:           stock.ID,
			IdStuff:      stock.IdStuff,
			Name:         stock.Name,
			Type:         stock.Type,
			Quantity:     stock.Quantity,
			Unit:         stock.Unit,
			Price:        stock.Price,
			Description:  stock.Description,
			CreatedAt:    stock.CreatedAt,
			UpdatedAt:    stock.UpdatedAt,
			DeletedAt:    stock.DeletedAt,
			InitialStock: 0,
			FinalStock:   0,
			InStock:      0,
			OutStock:     0,
		})
	}

	// group by month and year
	periodMap := make(map[string][]domain.StockReport)

	for _, stock := range finalPeriodStock {
		period := fmt.Sprintf("%d-%d", stock.CreatedAt.Month(), stock.CreatedAt.Year())
		periodMap[period] = append(periodMap[period], stock)
	}

	periodStocks := []domain.PeriodStock{}
	for period, finalPeriodStock := range periodMap {
		periodStocks = append(periodStocks, domain.PeriodStock{
			Date:   period,
			Stocks: finalPeriodStock,
		})
	}

	// barang masuk
	ins, err := c.insRepository.RetrieveIns()
	if err != nil {
		return nil, err
	}
	periodIns := []domain.PeriodIn{}
	for _, in := range ins {
		period := fmt.Sprintf("%d-%d", in.CreatedAt.Month(), in.CreatedAt.Year())
		periodIns = append(periodIns, domain.PeriodIn{
			Date: period,
			Ins:  ins,
		})
	}

	// barang keluar
	outs, err := c.outsRepository.RetrieveOuts()
	if err != nil {
		return nil, err
	}
	periodOuts := []domain.PeriodOut{}
	for _, out := range outs {
		period := fmt.Sprintf("%d-%d", out.CreatedAt.Month(), out.CreatedAt.Year())
		periodOuts = append(periodOuts, domain.PeriodOut{
			Date: period,
			Outs: outs,
		})
	}

	// retur barang
	returns, err := c.returnRepository.RetrieveRtrs()
	if err != nil {
		return nil, err
	}
	periodReturns := []domain.PeriodRtr{}
	for _, rtr := range returns {
		period := fmt.Sprintf("%d-%d", rtr.CreatedAt.Month(), rtr.CreatedAt.Year())
		periodReturns = append(periodReturns, domain.PeriodRtr{
			Date: period,
			Rtrs: returns,
		})
	}

	for _, periodStock := range periodStocks {
		for _, periodIn := range periodIns {
			if periodIn.Date == periodStock.Date {
				for c, periodOut := range periodOuts {
					if periodOut.Date == periodStock.Date {
						stockQuantity := periodStock.Stocks[c].Quantity
						inQuantity := periodIn.Ins[c].Quantity
						outQuantity := periodOut.Outs[c].Order.TotalOrder
						var returnQuantity float64
						if c >= len(periodReturns) {
							returnQuantity = 0
						} else {
							returnQuantity = periodReturns[c].Rtrs[c].TotalReturn
						}
						initialStock := stockQuantity - inQuantity - outQuantity
						if initialStock < 0 {
							initialStock = 0
						}

						finalStock := stockQuantity
						fmt.Println("final", finalStock, "in", inQuantity, "out", outQuantity, "return", returnQuantity, "stock", stockQuantity)
						if finalStock < 0 {
							finalStock = 0
						}

						periodStock.Stocks[c].InitialStock = initialStock
						periodStock.Stocks[c].FinalStock = finalStock
						periodStock.Stocks[c].InStock = inQuantity
						periodStock.Stocks[c].OutStock = outQuantity
					}
				}
			}
		}
	}

	return periodStocks, nil
}
