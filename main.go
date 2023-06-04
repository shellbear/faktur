package main

import (
	"log"
	"os"

	"github.com/shellbear/faktur/api"
)

const defaultPort = "8080"
const defaultHost = "127.0.0.1"

func main() {
	host := os.Getenv("FAKTUR_HOST")
	if host == "" {
		host = defaultHost
	}

	port := os.Getenv("FAKTUR_PORT")
	if port == "" {
		port = defaultPort
	}

	e, err := api.New()
	if err != nil {
		log.Fatal(err)
	}

	e.Logger.Fatal(e.Start(host + ":" + port))
}
