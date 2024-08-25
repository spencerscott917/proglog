package main

import (
	"log"

	"github.com/spencerscot917/proglog/internal/server"
	"github.com/spencerscott917/proglog/internal/server"
)

func main() {
	srv := server.NewHttpServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
