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
	if sum == 0 {
		return 0, nil
	}
	rate := float64(sum) / float64(count)
	return rate, nil
}
