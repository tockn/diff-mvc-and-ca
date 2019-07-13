package service

import (
	"context"

	"github.com/tockn/diff-mvc-and-ca/domain/repository"
)

type Item interface {
	CalcItemRateByID(ctx context.Context, id int64) (float64, error)
}

type item struct {
	itemRepo   repository.Item
	reviewRepo repository.Review
}

func NewItem(itemRepo repository.Item, reviewRepo repository.Review) Item {
	return &item{
		itemRepo:   itemRepo,
		reviewRepo: reviewRepo,
	}
}

func (i *item) CalcItemRateByID(ctx context.Context, id int64) (float64, error) {
	count, err := i.reviewRepo.CountByItemID(ctx, id)
	if err != nil {
		return 0, err
	}
	sum, err := i.reviewRepo.SumOfRateByItemID(ctx, id)
	if err != nil {
		return 0, err
	}
	if sum == 0 {
		return 0, nil
	}
	rate := float64(sum) / float64(count)
	return rate, nil
}

func (i *item) UpdateRanking() {
}
