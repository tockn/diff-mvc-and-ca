package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tockn/diff-mvc-and-ca/external/api/request"
	"github.com/tockn/diff-mvc-and-ca/usecase"
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
	req := request.GetItem{ID: c.Param("itemID")}

	ctx := c.Request.Context()
	ctx = addGinContext(ctx, c)

	ctr.it.GetItem(ctx, req.ToInput())
}

func (ctr *controller) PostItem(c *gin.Context) {
	var req request.PostItem
	_ = c.BindJSON(&req)

	ctx := c.Request.Context()
	ctx = addGinContext(ctx, c)

	ctr.it.PostItem(ctx, req.ToInput())
}

func (ctr *controller) GetReview(c *gin.Context) {
	req := request.GetReview{ID: c.Param("reviewID")}

	ctx := c.Request.Context()
	ctx = addGinContext(ctx, c)

	ctr.it.GetReview(ctx, req.ToInput())
}

func (ctr *controller) PostReview(c *gin.Context) {
	var req request.PostReview
	_ = c.BindJSON(&req)
	req.ItemID = c.Param("itemID")

	ctx := c.Request.Context()
	ctx = addGinContext(ctx, c)

	ctr.it.PostReview(ctx, req.ToInput())
}
