package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ItemJson struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Rate    float64 `json:"rate"`
	Ranking int64   `json:"ranking"`
}

type Item struct {
	ID        int64
	Name      string
	Rate      float64
	Reviews   []*Review `gorm:"foreignkey:ItemID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FindOneItem(db *gorm.DB, id int64) (*Item, error) {
	var item Item
	err := db.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// GetItemRankingは、引数に渡されたidに該当するitemのランキングを返す
func GetItemRanking(db *gorm.DB, id int64) (int64, error) {
	var item Item
	if err := db.Select("rate").First(&item, id).Error; err != nil {
		return 0, err
	}

	var rank int64
	row := db.DB().QueryRow(`
SELECT
	COUNT(*) + 1
FROM
	items
WHERE
	rate > ?`, item.Rate)

	if err := row.Scan(&rank); err != nil {
		return 0, err
	}

	return rank, nil
}

func (i *Item) Insert(db *gorm.DB) error {
	return db.Create(i).Error
}
