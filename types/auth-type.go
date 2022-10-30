package types

type FetchUserResponseType struct {
	Success bool
	User    UserType
}

type AuthUserResponseType struct {
	Success bool         `json:"id,omitempty"`
	User    AuthUserType `json:"user,omitempty"`
}

type AuthUserType struct {
	ID        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Full_Name string `json:"full_name,omitempty"`
	Telephone string `json:"telephone,omitempty"`
	Password  string `json:"password,omitempty"`
	Google    string `json:"google,omitempty"`
}
