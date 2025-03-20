package models

type LoginResponse struct {
	Token   string `json:"token,omitempty"`
	Message string `json:"message,omitempty"`
}
