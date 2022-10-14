package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"math"
	"net/http"
	"os"
	"time"
)

const WebPort = "8081"

type Config struct {
	Rabbit *amqp.Connection
}

func main() {
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()
	log.Println("listening for and consuming RabbitMQ message...")

	app := Config{Rabbit: rabbitConn}

	log.Printf("Starting broker service on %s", WebPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", WebPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	// try connect
	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672")
		if err != nil {
			fmt.Println("RabbitMQ not ready yet")
			counts++
		} else {
			connection = c
			break
		}
		// 1, 2, 4, 16
		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off...")
		time.Sleep(backOff)
	}

	return connection, nil
}
