package infrastructure

import "github.com/gin-gonic/gin"

func NewServer() *gin.Engine {
	return gin.Default()
}
