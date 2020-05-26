package models

import (
	"time"
)

// UserEvent User Events for Auditing
// GET /userEvent your user events
type UserEvent struct {
	ID             float64      `json:"id,omitempty"`
	Type           string       `json:"type"`
	Status         string       `json:"status"`
	UserID         float64      `json:"userId"`
	CreatedByID    float64      `json:"createdById"`
	IP             string       `json:"ip,omitempty"`
	GeoipCountry   string       `json:"geoipCountry,omitempty"`
	GeoipRegion    string       `json:"geoipRegion,omitempty"`
	GeoipSubRegion string       `json:"geoipSubRegion,omitempty"`
	EventMeta      *interface{} `json:"eventMeta,omitempty"`
	Created        time.Time    `json:"created"`
}
