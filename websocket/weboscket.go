package websocket

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
