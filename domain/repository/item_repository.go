package repository

import (
	"context"

	"github.com/tockn/diff-mvc-and-ca/domain/entity"
)

type Item interface {
	FindByID(ctx context.Context, id int64) (*entity.Item, error)
	Save(ctx context.Context, name string) (*entity.Item, error)
	GetRankingByID(ctx context.Context, id int64) (int64, error)
	UpdateRateByID(ctx context.Context, id int64, rate float64) (*entity.Item, error)
}
