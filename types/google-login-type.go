package types

import "time"

type SignUpUserCreatedType struct {
	Email     string `json:"email,omitempty"`
	Full_Name string `json:"full_name,omitempty"`
	Telephone string `json:"telephone,omitempty"`
	Google    string `json:"google,omitempty"`
}

type SignUpSuccesfulResponseType struct {
	Success   bool                  `json:"success"`
	User      SignUpUserCreatedType `json:"user"`
	Token     string                `json:"token"`
	ExpiresAt time.Time             `json:"expires_at"`
	Message   string                `json:"message"`
}
