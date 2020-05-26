package models

import (
	"time"
)

// Announcement Public Announcements
// GET /announcement site announcement
// GET /announcement/urgent urgent (banner) announcement
type Announcement struct {
	ID      float32   `json:"id"`
	Link    string    `json:"link,omitempty"`
	Title   string    `json:"title,omitempty"`
	Content string    `json:"content,omitempty"`
	Date    time.Time `json:"date,omitempty"`
}
