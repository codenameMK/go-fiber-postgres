package main

import (
	"log"
)

func main() {
	store, err := NewPostgressStore()
	if err != nil {
		log.Fatal(err)
	}
	if err:= store.Init(); err != nil {
		log.Fatal(err)
	}
	server := NewAPIServer(":4000" , store)
	server.Run()
}