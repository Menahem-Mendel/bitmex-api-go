package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/Menahem-Mendel/bitmex-api-go/websocket"
)

var key = "KBvdkF8OdLXmio5vJr0upDQC"
var secret = "N3nvdzkdhJNkpnFQVHJpa-vh2P6XhkVaMDmZE-AEzQMmKk5j"

var wg sync.WaitGroup

func main() {
	// f, _ := os.Create("m.trace")
	// trace.Start(f)
	// defer trace.Stop()

	c := websocket.New("www.bitmex.com", "realtime").Auth(key, secret)

	c.Connect()
	defer c.Disconnect()

	if err := c.Subscribe(websocket.Trade, ""); err != nil {
		log.Println("can't subscribe to trades", err)
		return
	}

	if err := c.Subscribe(websocket.OrderBookL2, ""); err != nil {
		log.Println("can't subscribe to order book", err)
		return
	}

	go func() {
		for v := range c.Listen() {
			switch v.(type) {
			case []byte:
				fmt.Printf("%s\n", v)
			default:
				fmt.Println(v)
			}

		}
	}()

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	<-interrupt
	fmt.Println("INTERRUPT")
}
