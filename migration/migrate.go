package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
	"github.com/tockn/diff-mvc-and-ca/model"
)

func main() {
	db, err := gorm.Open("mysql", "root:password@/diffmvca?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	db.CreateTable(&model.Item{})
	db.CreateTable(&model.Review{})

	item := model.Item{
		Name: "name1",
		Rate: 100,
	}
	db.Create(&item)

	review1 := model.Review{
		Rate:   10,
		ItemID: 1,
	}
	db.Create(&review1)

	review2 := model.Review{
		Rate:   20,
		ItemID: 1,
	}
	db.Create(&review2)

}
