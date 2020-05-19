package models

// ModelError
type ModelError struct {
	Error *Error `json:"error"`
}

// Error
type Error struct {
	Message string `json:"message,omitempty"`
	Name    string `json:"name,omitempty"`
}
