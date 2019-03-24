package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ReviewJson struct {
	ID     string  `json:"id"`
	Rate   float64 `json:"rate"`
	ItemID string  `json:"item_id"`
}

type Review struct {
	ID        int64
	Rate      float64
	ItemID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FindOneReview(db *gorm.DB, id int64) (*Review, error) {
	var review Review
	if err := db.First(&review, id).Error; err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *Review) Insert(db *gorm.DB) error {
	if err := db.Create(r).Error; err != nil {
		return err
	}

	var sum, count float64
	row := db.DB().QueryRow(`
SELECT
	SUM(rate), COUNT(*)
FROM
	reviews 
WHERE
	item_id = ?`, r.ItemID)

	if err := row.Scan(&sum, &count); err != nil {
		return err
	}

	rate := sum / count

	if err := db.Model(&Item{}).Update("rate", rate).Error; err != nil {
		return err
	}

	return nil
}
