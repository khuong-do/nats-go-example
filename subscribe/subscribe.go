package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"sync"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Use a WaitGroup to wait for a message to arrive
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Subscribe
	if _, err := nc.Subscribe("my-subscribe", func(m *nats.Msg) {
		fmt.Println(string(m.Data))
		wg.Done()
	}); err != nil {
		log.Fatal(err)
	}

	wg.Wait()

}
