package usecase

import (
	"context"
	"time"

	"discount_service/usecase/dto"

	"github.com/go-kit/kit/log"
)

type repo struct {
	db     map[string][]dto.Discount
	logger log.Logger
}

func NewRepo(logger log.Logger) Repository {
	return &repo{
		db:     map[string][]dto.Discount{},
		logger: logger,
	}
}

func (repo *repo) InsertDiscount(ctx context.Context, id string, amount int) error {
	discount := &dto.Discount{
		Amount:    amount,
		CreatedAt: time.Time(time.Now()),
	}

	repo.db[id] = append(repo.db[id], *discount)
	return nil
}

func (repo *repo) GetDiscountsByID(ctx context.Context, id string) ([]dto.Discount, error) {
	return repo.db[id], nil
}

func (repo *repo) GetDiscounts(ctx context.Context) (map[string][]dto.Discount, error) {
	return repo.db, nil
}
