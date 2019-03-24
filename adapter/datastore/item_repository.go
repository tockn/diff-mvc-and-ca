package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/tockn/diff-mvc-and-ca/adapter/datastore/mysql"
	"github.com/tockn/diff-mvc-and-ca/domain/entity"
	"github.com/tockn/diff-mvc-and-ca/domain/repository"
)

// Implement repository.Item
type item struct {
	db *gorm.DB
}

func NewItem(db *gorm.DB) repository.Item {
	return &item{
		db: db,
	}
}

func (i *item) FindByID(id int64) (*entity.Item, error) {
	var mItem mysql.Item
	if err := i.db.First(&mItem, id).Error; err != nil {
		return nil, err
	}
	return mItem.ToEntity(), nil
}

func (i *item) GetRankingByID(id int64) (int64, error) {
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

func (i *item) UpdateRateByID(id int64, rate float64) (*entity.Item, error) {
	var mItem mysql.Item
	if err := i.db.First(&mItem, id).Update("rate", rate).Error; err != nil {
		return nil, err
	}
	return mItem.ToEntity(), nil
}
