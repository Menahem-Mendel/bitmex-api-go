package models

import (
	"time"
)

// Chat Trollbox Data
// GET  /chat chat messages
// POST /chat Send a chat message
type Chat struct {
	ID        float32   `json:"id,omitempty"`
	Date      time.Time `json:"date"`
	User      string    `json:"user"`
	Message   string    `json:"message"`
	HTML      string    `json:"html"`
	FromBot   bool      `json:"fromBot,omitempty"`
	ChannelID float64   `json:"channelID,omitempty"`
}

// ChatChannel Trollbox Data
// GET /chat/channels available channels
type ChatChannel struct {
	ID   float32 `json:"id,omitempty"`
	Name string  `json:"name"`
}

// ConnectedUsers Trollbox Data
// GET /chat/connected connected users
type ConnectedUsers struct {
	Users float32 `json:"users,omitempty"`
	Bots  float32 `json:"bots,omitempty"`
}
