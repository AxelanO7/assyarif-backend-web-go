package usecase

import (
	"assyarif-backend-web-go/domain"
	"context"
	"fmt"
	"time"
)

type rtrUseCase struct {
	rtrRepository  domain.RtrRepository
	contextTimeout time.Duration
}

func NewRtrUseCase(rtr domain.RtrRepository, t time.Duration) domain.RtrUseCase {
	return &rtrUseCase{
		rtrRepository:  rtr,
		contextTimeout: t,
	}
}

func (c *rtrUseCase) FetchRtrByID(ctx context.Context, id uint) (*domain.Rtr, error) {
	res, err := c.rtrRepository.RetrieveRtrByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *rtrUseCase) FetchRtrs(ctx context.Context) ([]domain.Rtr, error) {
	res, err := c.rtrRepository.RetrieveRtrs()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *rtrUseCase) CreateRtr(ctx context.Context, req *domain.Rtr) (*domain.Rtr, error) {
	res, err := c.rtrRepository.CreateRtr(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *rtrUseCase) UpdateRtr(ctx context.Context, req *domain.Rtr) (*domain.Rtr, error) {
	res, err := c.rtrRepository.UpdateRtr(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *rtrUseCase) DeleteRtr(ctx context.Context, id uint) error {
	err := c.rtrRepository.DeleteRtr(id)
	if err != nil {
		return err
	}
	return nil
}

func (c *rtrUseCase) GetRtrsByPeriod(ctx context.Context) ([]domain.PeriodRtr, error) {
	// get rtrs
	rtrs, err := c.rtrRepository.RetrieveRtrs()
	if err != nil {
		return nil, err
	}

	// group by month and year
	periodMap := make(map[string][]domain.Rtr)

	for _, rtr := range rtrs {
		period := fmt.Sprintf("%d-%d", rtr.CreatedAt.Month(), rtr.CreatedAt.Year())
		periodMap[period] = append(periodMap[period], rtr)
	}

	periodRtrs := []domain.PeriodRtr{}
	for period, rtrs := range periodMap {
		periodRtrs = append(periodRtrs, domain.PeriodRtr{
			Date: period,
			Rtrs: rtrs,
		})
	}

	return periodRtrs, nil
}
