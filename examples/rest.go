package main

import (
	"context"
	"fmt"
	"log"
	"time"

	bitmex "github.com/Menahem-Mendel/bitmex-api-go/rest"
)

var key = "KBvdkF8OdLXmio5vJr0upDQC"
var secret = "N3nvdzkdhJNkpnFQVHJpa-vh2P6XhkVaMDmZE-AEzQMmKk5j"

func main() {
	// trades()
	tradeBins()
}

func trades() {
	c, err := bitmex.NewClient(false).Auth(key, time.Now().Add(time.Hour).Unix())
	if err != nil {
		fmt.Println(err)
	}
	ctx := context.WithValue(context.Background(), bitmex.KeyCtxSecret, secret)

	f := bitmex.TradeConf{
		Symbol:  bitmex.XBTUSD,
		Count:   bitmex.MAXCount,
		EndTime: time.Now(),
	}
	t, err := c.GetTrades(ctx, f)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range t {
		fmt.Println(v)
	}
}

func tradeBins() {
	c, err := bitmex.NewClient(false).Auth(key, time.Now().Add(time.Hour).Unix())
	if err != nil {
		fmt.Println(err)
	}

	ctx := context.WithValue(context.Background(), bitmex.KeyCtxSecret, secret)

	f := bitmex.TradeBucketedConf{
		Symbol:  bitmex.XBTUSD,
		Count:   bitmex.MAXCount,
		BinSize: bitmex.Minute,
		EndTime: time.Now(),
	}

	t, err := c.GetTradeBucketeds(ctx, f)
	if err != nil {
		log.Fatal(err)
	}

	for i, v := range t {
		fmt.Printf("%d: %v\n", i, v)
	}
}
