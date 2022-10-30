package controller

import (
	"context"
	"frontendmod/config"
	"frontendmod/types"
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
		templates.ExecuteTemplate(w, "login.html", types.LoginDataType{Error: "Could not authenticate with Google account."})
		return
	}

	code := r.URL.Query()["code"][0]

	googleConfig := config.SetupConfig()

	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Println("Code-Token exchange failed.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		templates.ExecuteTemplate(w, "login.html", types.LoginDataType{Error: "Could not authenticate with Google account."})
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		log.Println("User data fetch failed.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		templates.ExecuteTemplate(w, "login.html", types.LoginDataType{Error: "Could not authenticate with Google account."})
		return
	}

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Json parsing failed.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		templates.ExecuteTemplate(w, "login.html", types.LoginDataType{Error: "Could not authenticate with Google account."})
		return
	}

	_, response := HandleGoogleLogin(w, userData)
	log.Println("HandleGoogleLogin response:", response)
	if response.Success {
		http.Redirect(w, r, "/edit-profile", http.StatusSeeOther)
		templates.ExecuteTemplate(w, "edit-profile.html", response.User)
		return
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		templates.ExecuteTemplate(w, "login.html", types.LoginDataType{Error: "Could not authenticate with Google account."})
		return
	}
}
