package main

import (
	"github.com/igorsoto/go/books/distributed-services-with-go/proglog/internal/server"
	"log"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
