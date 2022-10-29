package controller

import "time"

type FetchUserResponseType struct {
	Success bool
	User    UserType
}

type UserType struct {
	ID        string
	Email     string
	Full_Name string
	Telephone string
	Google    string
}

type AuthenticatedUserType struct {
	Success bool
	User    UserType
	Token   string
	Message string
}

type EditProfileResponseType struct {
	Success           bool
	Message           string
	Full_Name_Updated bool
	Telephone_Updated bool
	Email_Updated     bool
}

type LoginDataType struct {
	Error     string
	GoogleUrl string
}

type AuthUserType struct {
	ID        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Full_Name string `json:"full_name,omitempty"`
	Telephone string `json:"telephone,omitempty"`
	Password  string `json:"password,omitempty"`
	Google    string `json:"google,omitempty"`
}

type GoogleCallbackDataType struct {
	ID             string `json:"id,omitempty"`
	Email          string `json:"email,omitempty"`
	Verified_Email bool   `json:"verified_email,omitempty"`
	Name           string `json:"name,omitempty"`
	Given_Name     string `json:"given_name,omitempty"`
	Family_Name    string `json:"family_name,omitempty"`
	Picture        string `json:"picture,omitempty"`
	Locale         string `json:"locale,omitempty"`
	Message        string `json:"success,omitempty"`
	Success        string `json:"message,omitempty"`
}

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

type GoogleCallbackFailedResponse struct {
	Success bool   `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
}

type Successin struct {
	Success   bool   `json:"success,omitempty"`
	Message   string `json:"message,omitempty"`
	ID        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Full_Name string `json:"full_name,omitempty"`
}
