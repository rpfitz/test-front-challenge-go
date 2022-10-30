package middleware

import (
	"frontendmod/controller"
	"frontendmod/service"
	"frontendmod/types"
	"net/http"
)

func IsUserAuthenticated(w http.ResponseWriter, r *http.Request) (types.AuthUserResponseType, bool) {
	var failedResponse types.AuthUserResponseType

	authorizationCookie, err := controller.GetCookie(r, "Authorization")
	if err != nil {
		return failedResponse, false
	}

	_, response, err := service.GETAuthUser2(authorizationCookie)

	if response.Success == true {
		return response, true
	}

	return failedResponse, false
}
