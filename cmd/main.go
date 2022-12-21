package main

import (
	"log"

	"github.com/arimatakao/bankapi/api"
	"github.com/arimatakao/bankapi/storage"
)

func main() {
	storage, err := storage.NewPostgresStorage("postgres", "accounts", "1234")
	if err != nil {
		log.Fatal(err)
	}

	s := api.NewAPIServer(":8080", storage)
	s.Run()
}
