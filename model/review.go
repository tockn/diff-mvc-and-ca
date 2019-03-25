package model

import (
	"errors"
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

	if !r.Validate() {
		return errors.New("validation error")
	}

	if err := db.Create(r).Error; err != nil {
		return err
	}

	rate, err := CalcRateByItemID(db, r.ItemID)
	if err != nil {
		return err
	}

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

func CalcRateByItemID(db *gorm.DB, itemID int64) (float64, error) {
	var sum, count float64
	row := db.DB().QueryRow(`
SELECT
	SUM(rate), COUNT(*)
FROM
	reviews 
WHERE
	item_id = ?`, itemID)

	if err := row.Scan(&sum, &count); err != nil {
		return 0, err
	}

	rate := sum / count
	return rate, nil
}
