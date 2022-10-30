package controller

import (
	"frontendmod/env"
	"frontendmod/service"
	"frontendmod/types"
	"log"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	statusCode, _ := service.MakePostRequest(r, map[string]interface{}{}, env.BACK_END_HOST_URL, env.BACK_END_LOGOUT_ENDPOINT)
	if statusCode != 200 {
		log.Println("LogoutController - Error expiring session on API.")
	}

	cookiesToRemove := []string{"Authorization", "Email", "Full_Name", "Telephone", "Authorization", "Google"}
	ClearCookies(w, cookiesToRemove)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
	templates.ExecuteTemplate(w, "login.html", types.LoginDataType{GoogleUrl: env.FRONT_END_GOOGLE_LOGIN_URL})
	return
}
