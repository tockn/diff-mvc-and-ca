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

	review.HID = idStr
	review.ItemHID = itemID
	c.JSON(http.StatusOK, review)
}

func (r *Review) New(c *gin.Context) {
	var review model.Review
	if err := c.BindJSON(&review); err != nil {
		return
	}

	review.ItemHID = c.Param("itemID")

	if err := review.Insert(r.db, r.hash); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, review)
}
