package models

// ModelError error message
type Error struct {
	Error *struct {
		Message string `json:"message,omitempty"`
		Name    string `json:"name,omitempty"`
	} `json:"error"`
}
