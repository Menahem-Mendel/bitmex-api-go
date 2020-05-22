package websocket

const (
// "announcement",        // Site announcements
// "chat",                // Trollbox chat
// "connected",           // Statistics of connected users/bots
// "funding",             // Updates of swap funding rates. Sent every funding interval (usually 8hrs)
// "instrument",          // Instrument updates including turnover and bid/ask
// "insurance",           // Daily Insurance Fund updates
// "liquidation",         // Liquidation orders as they're entered into the book
// "orderBookL2_25",      // Top 25 levels of level 2 order book
// "orderBookL2",         // Full level 2 order book
// "orderBook10",         // Top 10 levels using traditional full book push
// "publicNotifications", // System-wide notifications (used for short-lived messages)
// "quote",               // Top level of the book
// "quoteBin1m",          // 1-minute quote bins
// "quoteBin5m",          // 5-minute quote bins
// "quoteBin1h",          // 1-hour quote bins
// "quoteBin1d",          // 1-day quote bins
// "settlement",          // Settlements
// "trade",               // Live trades
// "tradeBin1m",          // 1-minute trade bins
// "tradeBin5m",          // 5-minute trade bins
// "tradeBin1h",          // 1-hour trade bins
// "tradeBin1d",          // 1-day trade bins

// "affiliate",   // Affiliate status, such as total referred users & payout %
// "execution",   // Individual executions; can be multiple per order
// "order",       // Live updates on your orders
// "margin",      // Updates on your current account balance and margin requirements
// "position",    // Updates on your positions
// "privateNotifications", // Individual notifications - currently not used
// "transact"     // Deposit/Withdrawal updates
// "wallet"       // Bitcoin address balance data, including total deposits & withdrawals
)

// import (
// 	"crypto/tls"
// 	"fmt"
// 	"log"
// 	"net/url"

// 	"github.com/gorilla/websocket"
// 	bitmex "github.com/Menahem-Mendel/bitmex-api-go"
// )

// var (
// 	Bitmex = "wss://www.bitmex/realtime"
// )

// type Client struct {
// 	addr string
// 	key string
// 	secret string
// 	socket
// }

// func New(key, secret, addr string) (*Client, error) {
// 	return &Client{
// 		addr: addr,
// 		key: key,
// 		secret: secret,
// 	}, nil
// }

// func (c *Client) SubscribeTrades(symbol string) (<-chan []byte, error) {
// 	var m = bitmex.{
// 		Op: "subscribe",
// 		Args: []interface{}{
// 			"orderBookL2_25:XBT",
// 			"trade:XBT",
// 		},
// 	}

// 	ws.WriteJSON(m)
// 	if err := ws.WriteJSON(m); err != nil {
// 		log.Printf("error writing json to connection: %v", err)
// 	}
// 	return nil
// }

// type socket struct {
// 	conn          *websocket.Conn
// 	TLSSkipVerify bool
// 	stream chan []byte
// }

// func (ws *socket) Connect() error {
// 	if ws != nil {
// 		return nil
// 	}

// 	d := websocket.Dialer{
// 		ReadBufferSize:  1024,
// 		WriteBufferSize: 1024,
// 		TLSClientConfig: &tls.Config{InsecureSkipVerify: ws.TLSSkipVerify},
// 	}

// 	log.Printf("connecting to websocket...")
// 	conn, resp, err := d.Dial(ws.BaseURL, nil)
// 	if err != nil {
// 		return fmt.Errorf("can't dial to %q, status code %d: %v", ws.BaseURL, resp.StatusCode, err)
// 	}

// 	ws.conn = conn

// 	conn, _, err := websocket.DefaultDialer.Dial("wss://www.bitmex.com/realtime", nil)
// 	if err != nil {
// 		log.Panicf("error dialing: %v", err)
// 	}

// 	w := &websock{}

// 	w.Conn = conn

// 	return nil
// }

// func (ws *socket) Disconnect() error {
// 	if err := ws.conn.Close(); err != nil {
// 		fmt.Errorf("error while disconnecting from the server: %v", err)
// 	}
// 	close(ws.Stream)
// 	return nil
// }

// func (ws *socket) Send(msg interface{}) error {
// 	if err := ws.conn.WriteJSON(msg); err != nil {
// 		return fmt.Errorf("error sending to the server: %v", err)
// 	}

// 	return nil
// }
