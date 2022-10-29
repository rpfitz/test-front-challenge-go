package controller

import (
	"context"
	"encoding/json"
	"frontendmod/config"
	"frontendmod/env"
	"io"
	"log"
	"net/http"
)

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	googleConfig := config.SetupConfig()
	url := googleConfig.AuthCodeURL("randomstate")
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query()["state"][0]
	if state != "randomstate" {
		log.Println("States doesn't match.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		templates.ExecuteTemplate(w, "login.html", LoginDataType{Error: "Could not authenticate with Google account."})
		return
	}

	code := r.URL.Query()["code"][0]

	googleConfig := config.SetupConfig()

	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Println("Code-Token exchange failed.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		templates.ExecuteTemplate(w, "login.html", LoginDataType{Error: "Could not authenticate with Google account."})
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		log.Println("User data fetch failed.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		templates.ExecuteTemplate(w, "login.html", LoginDataType{Error: "Could not authenticate with Google account."})
		return
	}

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Json parsing failed.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		templates.ExecuteTemplate(w, "login.html", LoginDataType{Error: "Could not authenticate with Google account."})
		return
	}

	_, response := handleGoogleLogin(userData)
	if response.Success {
		http.SetCookie(w, &http.Cookie{Name: "Email", Value: response.User.Email, HttpOnly: false, Path: "/"})
		http.SetCookie(w, &http.Cookie{Name: "Google", Value: response.User.Google, HttpOnly: false, Path: "/"})
		http.SetCookie(w, &http.Cookie{Name: "Authorization", Value: response.Token, HttpOnly: false, Path: "/"})
		http.SetCookie(w, &http.Cookie{Name: "Full_Name", Value: response.User.Full_Name, HttpOnly: false, Path: "/"})
		http.SetCookie(w, &http.Cookie{Name: "Telephone", Value: response.User.Telephone, HttpOnly: false, Path: "/"})
		http.Redirect(w, r, "/edit-profile", http.StatusSeeOther)
		templates.ExecuteTemplate(w, "edit-profile.html", AuthUserType{
			Email:     response.User.Email,
			Google:    response.User.Google,
			Full_Name: response.User.Full_Name,
			Telephone: response.User.Telephone,
		})
		return
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		templates.ExecuteTemplate(w, "login.html", LoginDataType{Error: "Could not authenticate with Google account."})
		return
	}
}

func handleGoogleLogin(callBackData []byte) (int, SignUpSuccesfulResponseType) {
	var response SignUpSuccesfulResponseType

	responseStatusCode, res := BytePOSTRequest(callBackData, env.BACK_END_HOST_URL, env.BACK_END_GOOGLE_AUTH_ENDPOINT)

	json.Unmarshal(res, &response)
	return responseStatusCode, response
}
