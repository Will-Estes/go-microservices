package main

import (
	"fmt"
	"log"
	"net/http"
)

const WebPort = "8081"

type Config struct {
}

func main() {
	app := Config{}

	log.Printf("Starting broker service on %s", WebPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", WebPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
