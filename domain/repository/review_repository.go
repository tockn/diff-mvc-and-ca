package repository

import (
	"context"

	"github.com/tockn/diff-mvc-and-ca/domain/entity"
)

type Review interface {
	FindByID(ctx context.Context, id int64) (*entity.Review, error)
	Save(ctx context.Context, rate float64, itemID int64) (*entity.Review, error)
	SumOfRateByItemID(ctx context.Context, id int64) (int64, error)
	CountByItemID(ctx context.Context, id int64) (int64, error)
}
