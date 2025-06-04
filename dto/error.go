package dto

type Error struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}
