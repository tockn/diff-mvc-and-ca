package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tockn/diff-mvc-and-ca/usecase"
	"github.com/tockn/diff-mvc-and-ca/usecase/input"
)

type Controller interface {
	GetItem(c *gin.Context)
	PostItem(c *gin.Context)
	GetReview(c *gin.Context)
	PostReview(c *gin.Context)
}

func NewController(it usecase.Interactor) Controller {
	return &controller{
		it: it,
	}
}

type controller struct {
	it usecase.Interactor
}

func (ctr *controller) GetItem(c *gin.Context) {
	var ipt input.GetItem
	ipt.ID = c.Param("itemID")

	ctx := c.Request.Context()
	ctx = addGinContext(ctx, c)

	ctr.it.GetItem(ctx, &ipt)
}

func (ctr *controller) PostItem(c *gin.Context) {
	var ipt input.PostItem
	_ = c.BindJSON(&ipt)

	ctx := c.Request.Context()
	ctx = addGinContext(ctx, c)

	ctr.it.PostItem(ctx, &ipt)
}

func (ctr *controller) GetReview(c *gin.Context) {
	var ipt input.GetReview
	ipt.ID = c.Param("reviewID")

	ctx := c.Request.Context()
	ctx = addGinContext(ctx, c)

	ctr.it.GetReview(ctx, &ipt)
}

func (ctr *controller) PostReview(c *gin.Context) {
	var ipt input.PostReview
	_ = c.BindJSON(&ipt)
	ipt.ItemID = c.Param("itemID")

	ctx := c.Request.Context()
	ctx = addGinContext(ctx, c)

	ctr.it.PostReview(ctx, &ipt)
}
