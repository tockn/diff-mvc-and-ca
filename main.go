package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/speps/go-hashids"

	"github.com/tockn/diff-mvc-and-ca/external/api"

	"github.com/tockn/diff-mvc-and-ca/registry"
)

func main() {
	hash, err := newHash()
	if err != nil {
		log.Fatal(err)
	}

	db, err := openDB()
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(os.Stdout, "diffmvca", log.Ltime)

	r := registry.NewRegistry(db, hash, logger)
	controller := r.NewController()
	server := newServer()
	api.NewRouter(server, controller)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func newHash() (*hashids.HashID, error) {
	h := hashids.NewData()
	h.MinLength = 5
	return hashids.NewWithData(h)
}

func openDB() (*gorm.DB, error) {
	return gorm.Open("mysql", "root:password@/diffmvca?charset=utf8mb4&parseTime=true")
}

func newServer() *gin.Engine {
	return gin.Default()
}
