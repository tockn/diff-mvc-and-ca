package datastore

import (
	"context"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/tockn/diff-mvc-and-ca/adapter/datastore/mysql"
	"github.com/tockn/diff-mvc-and-ca/domain/entity"
	"github.com/tockn/diff-mvc-and-ca/domain/repository"
)

// Implement repository.Item
type item struct {
	db      *gorm.DB
	redisDB *redis.Client
}

func NewItem(db *gorm.DB, red *redis.Client) repository.Item {
	return &item{
		db:      db,
		redisDB: red,
	}
}

func (i *item) FindByID(ctx context.Context, id int64) (*entity.Item, error) {
	var mItem mysql.Item
	if err := i.db.First(&mItem, id).Error; err != nil {
		return nil, err
	}
	return mItem.ToEntity(), nil
}

func (i *item) GetRankingByID(ctx context.Context, id int64) (int64, error) {
	var mItem mysql.Item
	if err := i.db.Select("rate").First(&mItem, id).Error; err != nil {
		return 0, err
	}

	var rank int64
	row := i.db.DB().QueryRow(`
SELECT
	COUNT(*) + 1
FROM
	items
WHERE
	rate > ?`, mItem.Rate)

	if err := row.Scan(&rank); err != nil {
		return 0, err
	}

	return rank, nil
}

func (i *item) Save(ctx context.Context, name string) (*entity.Item, error) {
	mItem := mysql.Item{
		Name: name,
		Rate: 0,
	}
	if err := i.db.Create(&mItem).Error; err != nil {
		return nil, err
	}
	return mItem.ToEntity(), nil
}

func (i *item) UpdateRateByID(ctx context.Context, id int64, rate float64) (*entity.Item, error) {
	var mItem mysql.Item
	if err := i.db.First(&mItem, id).Update("rate", rate).Error; err != nil {
		return nil, err
	}
	return mItem.ToEntity(), nil
}
