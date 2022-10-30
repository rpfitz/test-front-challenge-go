package types

type AuthenticatedUserType struct {
	Success bool
	User    UserType
	Token   string
	Message string
}
