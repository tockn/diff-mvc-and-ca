package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/speps/go-hashids"
	"github.com/tockn/diff-mvc-and-ca/model"
)

type Review struct {
	db   *gorm.DB
	hash *hashids.HashID
}

func NewReview(db *gorm.DB, hash *hashids.HashID) *Review {
	return &Review{
		db:   db,
		hash: hash,
	}
}

func (r *Review) Show(c *gin.Context) {
	idStr := c.Param("reviewID")
	id, err := model.DecodeID(r.hash, idStr)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	review, err := model.FindOneReview(r.db, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	itemID, err := model.EncodeID(r.hash, review.ItemID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	res := &model.ReviewJson{
		ID:     idStr,
		Rate:   review.Rate,
		ItemID: itemID,
	}
	c.JSON(http.StatusOK, res)
}

func (r *Review) New(c *gin.Context) {
	var reviewJson model.ReviewJson
	if err := c.BindJSON(&reviewJson); err != nil {
		return
	}

	itemID, err := model.DecodeID(r.hash, c.Param("itemID"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	review := model.Review{
		Rate:   reviewJson.Rate,
		ItemID: itemID,
	}
	if err := review.Insert(r.db); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	idStr, err := model.EncodeID(r.hash, review.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	res := &model.ReviewJson{
		ID:   idStr,
		Rate: review.Rate,
	}

	c.JSON(http.StatusCreated, res)

}
