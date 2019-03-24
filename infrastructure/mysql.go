package infrastructure

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func OpenDB() (*gorm.DB, error) {
	return gorm.Open("mysql", "root:password@/diffmvca?charset=utf8mb4&parseTime=true")
}
