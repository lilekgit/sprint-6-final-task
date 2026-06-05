package main

import (
	"log"
	"os"

	"sprint6finaltask/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.LstdFlags)
	srv := server.NewServer(logger)
	logger.Fatal(srv.HTTP.ListenAndServe())
}
