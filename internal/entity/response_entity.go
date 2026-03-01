package entity

type Response struct {
	Message   string `json:"message,omitempty"`
	Data      any    `json:"data,omitempty"`
	IsSuccess bool   `json:"isSuccess,omitempty"`
}
