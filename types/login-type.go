package types

type SuccessLoginResponseType struct {
	Success bool
	User    UserType
	Token   string
	Message string
}

type ErrorLoginType struct {
	Error     string
	GoogleUrl string
}

type LoginDataType struct {
	Error     string
	GoogleUrl string
}
