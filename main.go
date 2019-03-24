package main

import (
	"log"
	"os"

	"github.com/tockn/diff-mvc-and-ca/external/api"

	"github.com/tockn/diff-mvc-and-ca/infrastructure"
	"github.com/tockn/diff-mvc-and-ca/registry"
)

func main() {
	hash, err := infrastructure.NewHash()
	if err != nil {
		log.Fatal(err)
	}

	db, err := infrastructure.OpenDB()
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(os.Stdout, "diffmvca", log.Ltime)

	r := registry.NewRegistry(db, hash, logger)
	controller := r.NewController()
	server := infrastructure.NewServer()
	api.NewRouter(server, controller)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
