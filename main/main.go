package main

import (
	"github.com/specspace/plasma"
	"log"
)

func main() {
	srv, err := plasma.NewServerWithDefaults()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(srv.ListenAndServe())
}
