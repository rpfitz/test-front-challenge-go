package controller

import (
	"frontendmod/env"
	"log"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	statusCode, _ := MakePostRequest(r, map[string]interface{}{}, env.BACK_END_HOST_URL, env.BACK_END_LOGOUT_ENDPOINT)
	if statusCode != 200 {
		log.Println("LogoutController - Error expiring session on API.")
	}

	cookiesToRemove := []string{"Authorization", "Email", "Full_Name", "Telephone", "Authorization", "Google"}
	for i := range cookiesToRemove {
		clearCookies(w, cookiesToRemove[i])
	}

	googleLoginUrl := env.FRONT_END_GOOGLE_LOGIN_URL
	http.Redirect(w, r, "/login", http.StatusSeeOther)
	templates.ExecuteTemplate(w, "login.html", LoginDataType{GoogleUrl: googleLoginUrl})
	return
}

func clearCookies(w http.ResponseWriter, cookieName string) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})
}
