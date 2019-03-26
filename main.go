package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/speps/go-hashids"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tockn/diff-mvc-and-ca/controller"
)

func main() {

	db, err := gorm.Open("mysql", "root:password@/diffmvca?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	h := hashids.NewData()
	h.MinLength = 5
	hashID, err := hashids.NewWithData(h)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(hashID.EncodeInt64([]int64{1}))

	r := gin.Default()

	item := controller.NewItem(db, hashID)
	r.GET("/items/:itemID", item.Show)
	r.POST("/items", item.New)

	review := controller.NewReview(db, hashID)
	r.GET("/items/:itemID/reviews/:reviewID", review.Show)
	r.POST("/items/:itemID/reviews", review.New)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
