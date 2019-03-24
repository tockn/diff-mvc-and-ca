package repository

import "github.com/tockn/diff-mvc-and-ca/domain/entity"

type Item interface {
	FindByID(id int64) (*entity.Item, error)
	GetRankingByID(id int64) (int64, error)
	UpdateRateByID(id int64, rate float64) (*entity.Item, error)
}
