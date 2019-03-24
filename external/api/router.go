package api

import "github.com/gin-gonic/gin"

func NewRouter(r *gin.Engine, c Controller) {
	r.GET("/items/:itemID", c.GetItem)

	r.GET("/items/:itemID/reviews/:reviewID", c.GetReview)
	r.POST("/items/:itemID/reviews", c.PostReview)
}
