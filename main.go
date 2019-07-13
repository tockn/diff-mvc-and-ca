package main

import (
	"log"
	"os"

	"github.com/go-redis/redis"

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

	red, err := openRedis()
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(os.Stdout, "diffmvca", log.Ltime)

	r := registry.NewRegistry(db, hash, logger, red)
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

func openRedis() (*redis.Client, error) {
	redisDB := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	_, err := redisDB.Ping().Result()
	return redisDB, err
}
