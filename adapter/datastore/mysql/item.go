package mysql

import (
	"time"

	"github.com/tockn/diff-mvc-and-ca/domain/entity"
)

type Item struct {
	ID        int64
	Name      string
	Rate      float64
	Reviews   []*Review `gorm:"foreignkey:ItemID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewItemFromEntity(ei *entity.Item) *Item {
	return &Item{
		ID:   ei.ID,
		Name: ei.Name,
		Rate: ei.Rate,
	}
}

func (i *Item) ToEntity() *entity.Item {
	return &entity.Item{
		ID:   i.ID,
		Name: i.Name,
		Rate: i.Rate,
	}
}
