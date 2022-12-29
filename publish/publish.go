package main

import (
	"github.com/nats-io/nats.go"
	"log"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	if err := nc.Publish("my-subscribe", []byte("My name is Khuong Do. Nice to meet you!")); err != nil {
		log.Fatal(err)
	}
}
