package datastore

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/tockn/diff-mvc-and-ca/adapter/datastore/mysql"
	"github.com/tockn/diff-mvc-and-ca/domain/entity"
	"github.com/tockn/diff-mvc-and-ca/domain/repository"
)

type review struct {
	db *gorm.DB
}

func NewReview(db *gorm.DB) repository.Review {
	return &review{
		db: db,
	}
}

func (r *review) FindByID(ctx context.Context, id int64) (*entity.Review, error) {
	var mReview mysql.Review
	if err := r.db.First(&mReview, id).Error; err != nil {
		return nil, err
	}
	return mReview.ToEntity(), nil
}

func (r *review) Save(ctx context.Context, rate float64, itemID int64) (*entity.Review, error) {
	mReview := mysql.Review{
		Rate:   rate,
		ItemID: itemID,
	}
	if err := r.db.Create(&mReview).Error; err != nil {
		return nil, err
	}
	return mReview.ToEntity(), nil
}

func (r *review) SumOfRateByItemID(ctx context.Context, id int64) (int64, error) {
	var sum int64
	row := r.db.DB().QueryRow(`
SELECT
	SUM(rate)
FROM
	reviews
WHERE
	item_id = ?`, id)
	if err := row.Scan(&sum); err != nil {
		return 0, err
	}
	return sum, nil
}

func (r *review) CountByItemID(ctx context.Context, id int64) (int64, error) {
	var count int64
	row := r.db.DB().QueryRow(`
SELECT
	COUNT(*)
FROM
	reviews
WHERE
	item_id = ?`, id)
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}
