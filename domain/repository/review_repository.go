package repository

import "github.com/tockn/diff-mvc-and-ca/domain/entity"

type Review interface {
	FindByID(id int64) (*entity.Review, error)
	Save(rate float64, itemID int64) (*entity.Review, error)
	SumOfRateByItemID(id int64) (int64, error)
	CountByItemID(id int64) (int64, error)
}
