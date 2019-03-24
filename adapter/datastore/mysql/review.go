package mysql

import (
	"time"

	"github.com/tockn/diff-mvc-and-ca/domain/entity"
)

type Review struct {
	ID        int64
	Rate      float64
	ItemID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewReviewFromEntity(er *entity.Review) *Review {
	return &Review{
		ID:   er.ID,
		Rate: er.Rate,
	}
}

func (r *Review) ToEntity() *entity.Review {
	return &entity.Review{
		ID:   r.ID,
		Rate: r.Rate,
	}
}
