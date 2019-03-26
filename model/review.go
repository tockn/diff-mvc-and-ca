package model

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Review struct {
	ID        int64     `json:"-"`
	HID       string    `json:"id" gorm:"-"`
	Rate      float64   `json:"rate"`
	ItemID    int64     `json:"-"`
	ItemHID   string    `json:"item_id" gorm:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func FindOneReview(db *gorm.DB, id int64) (*Review, error) {
	var review Review
	if err := db.First(&review, id).Error; err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *Review) Insert(db *gorm.DB) error {

	if !r.Validate() {
		return errors.New("validation error")
	}

	if err := db.Create(r).Error; err != nil {
		return err
	}

	var sum, count int64
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

	rate := CalculateRate(sum, count)

	if err := db.Model(&Item{}).Update("rate", rate).Error; err != nil {
		return err
	}

	return nil
}

func (r *Review) Validate() bool {
	if r.Rate < 1 || 5 < r.Rate {
		return false
	}
	return true
}

func CalculateRate(sum, count int64) float64 {
	if count == 0 {
		return 0
	}
	rate := float64(sum) / float64(count)
	return rate
}
