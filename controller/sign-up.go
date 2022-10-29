package controller

import (
	"log"
	"net/http"
)

func GetSignUp(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "sign-up.html", nil)
}

func PostSignUp(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	status, response := FetchPostAPISignUp(email, password)
	if status != "200 OK" {
		log.Println("SignUpGetController: Error fetching data.")
		return
	}

	if response.Success {
		http.SetCookie(w, &http.Cookie{Name: "Authorization", Value: response.Token, HttpOnly: false})
		http.SetCookie(w, &http.Cookie{Name: "Full_Name", Value: response.User.Full_Name, HttpOnly: false})
		http.SetCookie(w, &http.Cookie{Name: "Telephone", Value: response.User.Telephone, HttpOnly: false})
		http.SetCookie(w, &http.Cookie{Name: "Email", Value: response.User.Email, HttpOnly: false})
		http.SetCookie(w, &http.Cookie{Name: "Google", Value: response.User.Google, HttpOnly: false})
		http.Redirect(w, r, "/edit-profile", http.StatusSeeOther)
		return
	} else {
		templates.ExecuteTemplate(w, "sign-up.html", LoginDataType{Error: response.Message})
		return
	}
}
