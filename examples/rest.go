package main

import (
	"context"
	"fmt"
	"log"

	bitmex "github.com/Menahem-Mendel/bitmex-api-go"
	"github.com/Menahem-Mendel/bitmex-api-go/rest"
)

var key = "KBvdkF8OdLXmio5vJr0upDQC"
var secret = "N3nvdzkdhJNkpnFQVHJpa-vh2P6XhkVaMDmZE-AEzQMmKk5j"

func main() {
	// getOrders()
	// for i := 0; i < 100; i++ {
	// newOrder()
	// // }
	// getOrders()
	// deleteOrder()
	apiKey()
}

func apiKey() {
	c, err := rest.NewAuthClient(false, key)
	if err != nil {
		log.Println(err)
	}
	ctx := context.WithValue(context.Background(), bitmex.ContextAPIKey, secret)

	f := rest.APIKeyGetConf{
		Reverse: true,
	}

	out, err := c.GetAPIKey(ctx, f)
	if err != nil {
		log.Println(err)
	}

	for i, v := range out {
		fmt.Printf("%d: %v\n", i, v)
	}
}

func orderBookL2() {
	c, err := rest.NewAuthClient(false, key)
	if err != nil {
		log.Println(err)
	}
	ctx := context.WithValue(context.Background(), bitmex.ContextAPIKey, secret)

	f := rest.OrderBookL2Conf{
		Symbol: bitmex.XBTUSD,
		Depth:  0,
	}

	out, err := c.GetOrderBookL2(ctx, f)
	if err != nil {
		log.Println(err)
	}

	for i, v := range out {
		fmt.Printf("%d: %v\n", i, v)
	}
}

func deleteOrder() {
	c, err := rest.NewAuthClient(false, key)
	if err != nil {
		log.Println(err)
	}
	ctx := context.WithValue(context.Background(), bitmex.ContextAPIKey, secret)

	f := rest.OrderCancelConf{
		ClOrdID: "mendelVVUsfb2w",
	}

	out, err := c.CancelOrder(ctx, f)
	if err != nil {
		log.Println(err)
	}

	for i, v := range out {
		fmt.Printf("%d: %v\n", i, v)
	}
}

func getOrders() {
	c, err := rest.NewAuthClient(false, key)
	if err != nil {
		log.Println(err)
	}
	ctx := context.WithValue(context.Background(), bitmex.ContextAPIKey, secret)

	f := rest.OrderConf{
		Symbol: bitmex.XBTUSD,
	}

	out, err := c.GetOrders(ctx, f)
	if err != nil {
		log.Println(err)
	}

	for i, v := range out {
		fmt.Printf("%d: %v\n", i, v)
	}
}

func newOrder() {
	c, err := rest.NewAuthClient(false, key)
	if err != nil {
		log.Println(err)
	}
	ctx := context.WithValue(context.Background(), bitmex.ContextAPIKey, secret)

	if err != nil {
		log.Println(err)
	}

	f := rest.OrderNewConf{
		Symbol:   bitmex.XBTUSD,
		Price:    500,
		ClOrdID:  "mendel",
		OrderQty: 1,
		OrdType:  bitmex.Limit,
	}

	out, err := c.NewOrder(ctx, f)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%v\n", out)
}
