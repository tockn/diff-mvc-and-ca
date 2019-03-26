package model

import (
	"errors"
	"time"

	"github.com/speps/go-hashids"

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

func (r *Review) Insert(db *gorm.DB, hash *hashids.HashID) error {

	// バリデーション
	if err := r.Validate(); err != nil {
		return err
	}

	// 入力されたレビューデータの商品ハッシュIDを数値IDへ変換
	itemID, err := DecodeID(hash, r.ItemHID)
	if err != nil {
		return err
	}

	r.ItemID = itemID

	// 入力されたレビューデータを永続化
	if err := db.Create(r).Error; err != nil {
		return err
	}

	// 永続化したレビューデータの数値IDをハッシュIDへ変換
	idStr, err := EncodeID(hash, r.ID)
	if err != nil {
		return err
	}

	r.HID = idStr

	// レビューされた商品のレート更新
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

func (r *Review) Validate() error {
	if r.Rate < 1 || 5 < r.Rate {
		return errors.New("[Review] Rate validate error")
	}
	return nil
}

func CalculateRate(sum, count int64) float64 {
	if count == 0 {
		return 0
	}
	rate := float64(sum) / float64(count)
	return rate
}
