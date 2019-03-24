package service

import "github.com/tockn/diff-mvc-and-ca/domain/repository"

type Item interface {
	CalcItemRateByID(id int64) (float64, error)
}

type item struct {
	reviewRepo repository.Review
}

func NewItem(reviewRepo repository.Review) Item {
	return &item{
		reviewRepo: reviewRepo,
	}
}

func (i *item) CalcItemRateByID(id int64) (float64, error) {
	count, err := i.reviewRepo.CountByItemID(id)
	if err != nil {
		return 0, err
	}
	sum, err := i.reviewRepo.SumOfRateByItemID(id)
	if err != nil {
		return 0, err
	}
	rate := float64(count) / float64(sum)
	return rate, nil
}
