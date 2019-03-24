package api

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tockn/diff-mvc-and-ca/usecase"
	"github.com/tockn/diff-mvc-and-ca/usecase/output"
)

type presenter struct {
	logger *log.Logger
}

func NewPresenter(logger *log.Logger) usecase.Presenter {
	return &presenter{
		logger: logger,
	}
}

func (p *presenter) ViewItem(ctx context.Context, item *output.Item) {
	defer deleteGinContext(ctx)
	c := getGinContext(ctx)
	p.JSON(c, http.StatusOK, item)
}

func (p *presenter) ViewReview(ctx context.Context, review *output.Review) {
	defer deleteGinContext(ctx)
	c := getGinContext(ctx)
	p.JSON(c, http.StatusOK, review)
}

func (p *presenter) ViewError(ctx context.Context, err error) {
	defer deleteGinContext(ctx)
	c := getGinContext(ctx)
	p.logger.Println(err)
	p.JSON(c, http.StatusInternalServerError, "")
}

func (p *presenter) JSON(c *gin.Context, code int, v interface{}) {
	c.JSON(code, v)
}
