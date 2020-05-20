package models

// ModelError error message
type ModelError struct {
	Error *Error `json:"error"`
}

// Error error message field error
type Error struct {
	Message string `json:"message,omitempty"`
	Name    string `json:"name,omitempty"`
}
