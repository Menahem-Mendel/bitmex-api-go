package main

import (
	"fmt"
	"log"
	"time"

	bitmex "github.com/Menahem-Mendel/bitmex-api-go"
	"github.com/Menahem-Mendel/bitmex-api-go/rest"
)

var key = "KBvdkF8OdLXmio5vJr0upDQC"
var secret = "N3nvdzkdhJNkpnFQVHJpa-vh2P6XhkVaMDmZE-AEzQMmKk5j"

// var wg sync.WaitGroup

func main() {
	// getOrders()
	// for i := 0; i < 100; i++ {
	// wg.Add(1)
	// go newOrder()
	// // // }
	// getOrders()
	// deleteOrder()
	// amendOrder()
	// apiKey()
	// for i := 0; i < 100; i++ {
	// 	getTrades()
	// }
	// wg.Wait()

	c := rest.NewClient().Auth(secret, key)

	f2 := rest.TradeFilter{
		Symbol:    bitmex.XBTUSD,
		StartTime: time.Now().Add(time.Minute * -1),
		EndTime:   time.Now(),
		Count:     1,
	}

	for i := 0; i < 100; i++ {
		bs2, err := c.Trades.Get(f2)
		if err != nil {
			log.Printf("ERROR - %v", err)
		}
		for i, v := range bs2 {
			fmt.Printf("%d - %v\n", i, v)
		}
	}
}

// func apiKey() {
// 	c := rest.NewClient()

// 	ctx := context.WithValue(context.TODO(), bitmex.ContextAPIKey, secret)
// 	f := rest.APIKeyGetConf{}
// 	k, err := c.APIKeys.Get(ctx, f)
// 	if err != nil {
// 		log.Printf("%v", err)
// 		return
// 	}

// 	fmt.Println(*k)

// }

// // func amendOrder() {
// // 	c := rest.NewClient().Credentials(key)

// // 	ctx := context.WithValue(context.TODO(), bitmex.ContextAPIKey, secret)
// // 	f := rest.OrderAmendConf{
// // 		// OrderID:  "dae1929b-6560-5b30-98fb-155df949e418",
// // 		OrderQty:    5,
// // 		OrigClOrdID: "men_7mTDkagj",
// // 	}

// // 	ts, err := c.Orders.Amend(ctx, f)
// // 	if err != nil {
// // 		log.Printf("%v", err)
// // 		return
// // 	}

// // 	log.Println(*ts)
// // 	// wg.Done()
// // }

// // func newOrder() {
// // 	c := rest.NewClient().Credentials(key)

// // 	ctx := context.WithValue(context.TODO(), bitmex.ContextAPIKey, secret)
// // 	f := rest.OrderNewConf{
// // 		Symbol:   bitmex.XBTUSD,
// // 		ClOrdID:  "men_",
// // 		OrderQty: 5,
// // 	}

// // 	ts, err := c.Orders.New(ctx, f)
// // 	if err != nil {
// // 		log.Printf("%v", err)
// // 		return
// // 	}

// // 	log.Println(*ts)
// // }

// // func getOrders() {
// // 	c := rest.NewClient().Credentials(key)

// // 	ctx := context.WithValue(context.TODO(), bitmex.ContextAPIKey, secret)
// // 	f := rest.OrderConf{}

// // 	ts, err := c.Orders.Get(ctx, f)
// // 	if err != nil {
// // 		log.Printf("%v", err)
// // 		return
// // 	}

// // 	fmt.Println()
// // 	for i, v := range *ts {
// // 		log.Println(i, v)

// // 	}
// // 	fmt.Println()

// // }

// // //
// func getTrades() {
// 	c := rest.NewClient()

// 	f := rest.TradeConf{
// 		Symbol:    bitmex.XBTUSD,
// 		Count:     10,
// 		StartTime: time.Now().Local().Add(time.Hour * -1),
// 		EndTime:   time.Now().Local(),
// 	}

// 	ts, err := c.Trades.Get(f)
// 	if err != nil {
// 		log.Printf("%v", err)
// 		return
// 	}

// 	// fmt.Println(ts)

// 	for i, v := range *ts {
// 		log.Println(i, v)
// 	}
// }
