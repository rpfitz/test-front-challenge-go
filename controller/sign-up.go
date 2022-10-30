package controller

import (
	"frontendmod/types"
	"log"
	"net/http"
)

func GetSignUp(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "sign-up.html", nil)
}

func PostSignUp(w http.ResponseWriter, r *http.Request) {
	response, error := HandlePostSignUpRequest(w, r)
	if error != nil {
		log.Println(error)
	}

	if response.Success {
		http.Redirect(w, r, "/edit-profile", http.StatusSeeOther)
		return
	} else {
		templates.ExecuteTemplate(w, "sign-up.html", types.LoginDataType{Error: response.Message})
		return
	}
}
