package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/speps/go-hashids"
	"github.com/tockn/diff-mvc-and-ca/model"
)

type Item struct {
	db     *gorm.DB
	hash   *hashids.HashID
	logger *log.Logger
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

	res := &model.ItemJson{
		ID:      idStr,
		Name:    item.Name,
		Rate:    item.Rate,
		Ranking: rank,
	}

	c.JSON(http.StatusOK, res)
}

func (i *Item) New(c *gin.Context) {
	var itemJson model.ItemJson
	if err := c.BindJSON(&itemJson); err != nil {
		return
	}

	item := model.Item{
		Name: itemJson.Name,
		Rate: 0,
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
	res := model.ItemJson{
		ID:      idStr,
		Name:    item.Name,
		Rate:    item.Rate,
		Ranking: rank,
	}
	c.JSON(http.StatusCreated, res)
}
