package controller

import (
	"frontendmod/env"
	"frontendmod/types"
	"log"
	"net/http"
)

func GetLogin(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "login.html", types.LoginDataType{GoogleUrl: env.FRONT_END_GOOGLE_LOGIN_URL})
	return
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	response, error := HandlePostLoginRequest(w, r)
	if error != nil {
		log.Println(error)
	}

	if response.Success {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return
	}

	templates.ExecuteTemplate(w, "login.html", types.LoginDataType{Error: response.Message, GoogleUrl: env.FRONT_END_GOOGLE_LOGIN_URL})
	return
}
