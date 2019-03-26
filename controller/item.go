package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/speps/go-hashids"
	"github.com/tockn/diff-mvc-and-ca/model"
)

type Item struct {
	db   *gorm.DB
	hash *hashids.HashID
}

func NewItem(db *gorm.DB, hash *hashids.HashID) *Item {
	return &Item{
		db:   db,
		hash: hash,
	}
}

func (i *Item) Show(c *gin.Context) {
	idStr := c.Param("itemID")

	id, err := model.DecodeID(i.hash, idStr)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	item, err := model.FindOneItem(i.db, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	rank, err := model.GetItemRanking(i.db, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	item.HID = idStr
	item.Ranking = rank

	c.JSON(http.StatusOK, item)
}

func (i *Item) New(c *gin.Context) {
	var item model.Item
	if err := c.BindJSON(&item); err != nil {
		return
	}

	if err := item.Insert(i.db); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	idStr, err := model.EncodeID(i.hash, item.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	rank, err := model.GetItemRanking(i.db, item.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	item.HID = idStr
	item.Ranking = rank
	c.JSON(http.StatusCreated, item)
}
