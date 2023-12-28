package main

import (
	"log"
	"go-fiber-postgres/providers"
)

func main() {
	store, err := NewPostgressStore()
	config, err := providers.GetConfig("search-playground-service-configuration.yml")
	if err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	if err:= store.Init(); err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(":4000" , store)
	server.Run()
}