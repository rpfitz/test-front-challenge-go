package controller

import (
	"frontendmod/types"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	_, email, full_name, telephone, isGoogle := GetCookiesSession(r)

	templates.ExecuteTemplate(w, "profile.html", types.AuthUserType{
		Email:     email,
		Full_Name: full_name,
		Telephone: telephone,
		Google:    isGoogle,
	})
	return
}
