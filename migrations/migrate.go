package main

import (
	"log"

	"github.com/tockn/diff-mvc-and-ca/adapter/datastore/mysql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("mysql", "root:password@/diffmvca?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	db.CreateTable(&mysql.Item{})
	db.CreateTable(&mysql.Review{})

	item := mysql.Item{
		Name: "name1",
		Rate: 100,
	}
	db.Create(&item)

	review1 := mysql.Review{
		Rate:   10,
		ItemID: 1,
	}
	db.Create(&review1)

	review2 := mysql.Review{
		Rate:   20,
		ItemID: 1,
	}
	db.Create(&review2)

}
