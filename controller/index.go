package controller

import (
	"frontendmod/types"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func ExecTemplates() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

func SetCacheHeaders(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}

func SetAllCookies(w http.ResponseWriter, response types.SuccessLoginResponseType) {
	http.SetCookie(w, &http.Cookie{Name: "Authorization", Value: response.Token, HttpOnly: false})
	http.SetCookie(w, &http.Cookie{Name: "Full_Name", Value: response.User.Full_Name, HttpOnly: false})
	http.SetCookie(w, &http.Cookie{Name: "Telephone", Value: response.User.Telephone, HttpOnly: false})
	http.SetCookie(w, &http.Cookie{Name: "Email", Value: response.User.Email, HttpOnly: false})
	http.SetCookie(w, &http.Cookie{Name: "Google", Value: response.User.Google, HttpOnly: false})
}

func SetAllCookiesSignUp(w http.ResponseWriter, response types.AuthenticatedUserType) {
	http.SetCookie(w, &http.Cookie{Name: "Authorization", Value: response.Token, HttpOnly: false})
	http.SetCookie(w, &http.Cookie{Name: "Full_Name", Value: response.User.Full_Name, HttpOnly: false})
	http.SetCookie(w, &http.Cookie{Name: "Telephone", Value: response.User.Telephone, HttpOnly: false})
	http.SetCookie(w, &http.Cookie{Name: "Email", Value: response.User.Email, HttpOnly: false})
	http.SetCookie(w, &http.Cookie{Name: "Google", Value: response.User.Google, HttpOnly: false})
}

func SetAllGoogleCookies(w http.ResponseWriter, response types.SignUpSuccesfulResponseType) {
	http.SetCookie(w, &http.Cookie{Name: "Authorization", Value: response.Token, HttpOnly: false, Path: "/"})
	http.SetCookie(w, &http.Cookie{Name: "Full_Name", Value: response.User.Full_Name, HttpOnly: false, Path: "/"})
	http.SetCookie(w, &http.Cookie{Name: "Telephone", Value: response.User.Telephone, HttpOnly: false, Path: "/"})
	http.SetCookie(w, &http.Cookie{Name: "Email", Value: response.User.Email, HttpOnly: false, Path: "/"})
	http.SetCookie(w, &http.Cookie{Name: "Google", Value: response.User.Google, HttpOnly: false, Path: "/"})
}

func CookiesUpdateEditProfile(w http.ResponseWriter, responseBody types.UpdateProfileSuccessType, email string, full_name string, telephone string, newToken string) {
	if responseBody.Email_Updated {
		http.SetCookie(w, &http.Cookie{Name: "Authorization", Value: newToken, HttpOnly: false})
		http.SetCookie(w, &http.Cookie{Name: "Email", Value: email, HttpOnly: false})
	}

	if responseBody.Full_Name_Updated {
		http.SetCookie(w, &http.Cookie{Name: "Full_Name", Value: full_name, HttpOnly: false})
	}

	if responseBody.Telephone_Updated {
		http.SetCookie(w, &http.Cookie{Name: "Telephone", Value: telephone, HttpOnly: false})
	}
}

func GetCookiesSession(r *http.Request) (string, string, string, string, string) {
	token, err := GetCookie(r, "Authorization")

	if err != nil {
		log.Println("Error getting cookie Authorization")
	}
	email, err := GetCookie(r, "Email")

	if err != nil {
		log.Println("Error getting cookie Email")
	}
	full_name, err := GetCookie(r, "Full_Name")

	if err != nil {
		log.Println("Error getting cookie Full_Name")
	}
	telephone, err := GetCookie(r, "Telephone")

	if err != nil {
		log.Println("Error getting cookie Telephone")
	}
	google, err := GetCookie(r, "Google")
	if err != nil {
		log.Println("Error getting cookie Google")
	}

	return token, email, full_name, telephone, google
}

func WhatToEditProfile(userCookieEmail string, email string, userCookieFullName string, full_name string, userCookieTelephone string, telephone string) (bool, bool, bool) {
	var (
		editEmail     = true
		editFullName  = true
		editTelephone = true
	)

	if email == userCookieEmail {
		editEmail = false
	}
	if full_name == userCookieFullName {
		editFullName = false
	}
	if telephone == userCookieTelephone {
		editTelephone = false
	}

	return editEmail, editFullName, editTelephone
}

func GetCookie(r *http.Request, cookieName string) (string, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func ClearCookie(w http.ResponseWriter, cookieName string) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})
}

func ClearCookies(w http.ResponseWriter, cookiesToRemove []string) {
	for i := range cookiesToRemove {
		ClearCookie(w, cookiesToRemove[i])
	}
}
