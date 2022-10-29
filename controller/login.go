package controller

import (
	"frontendmod/env"
	"log"
	"net/http"
)

func GetLogin(w http.ResponseWriter, r *http.Request) {
	googleLoginUrl := env.FRONT_END_GOOGLE_LOGIN_URL
	templates.ExecuteTemplate(w, "login.html", LoginDataType{GoogleUrl: googleLoginUrl})
	return
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	status, response := FetchPostAPILogin(email, password)
	if status != "200 OK" {
		log.Println("LoginGetController: Error fetching data.")
		return
	}

	if response.Success {
		http.SetCookie(w, &http.Cookie{Name: "Authorization", Value: response.Token, HttpOnly: false})
		http.SetCookie(w, &http.Cookie{Name: "Full_Name", Value: response.User.Full_Name, HttpOnly: false})
		http.SetCookie(w, &http.Cookie{Name: "Telephone", Value: response.User.Telephone, HttpOnly: false})
		http.SetCookie(w, &http.Cookie{Name: "Email", Value: response.User.Email, HttpOnly: false})
		http.SetCookie(w, &http.Cookie{Name: "Google", Value: response.User.Google, HttpOnly: false})
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return
	}

	googleLoginUrl := env.FRONT_END_GOOGLE_LOGIN_URL
	templates.ExecuteTemplate(w, "login.html", LoginDataType{Error: response.Message, GoogleUrl: googleLoginUrl})
	return
}
