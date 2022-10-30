package middleware

import (
	"frontendmod/controller"
	"frontendmod/env"
	"frontendmod/types"
	"html/template"
	"net/http"
)

var templates *template.Template

func ExecTemplates() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

func ValidateAuthentication(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controller.SetCacheHeaders(w)

		_, isUserAuthenticated := IsUserAuthenticated(w, r)

		if isUserAuthenticated {
			if r.URL.Path == "/" || r.URL.Path == "/login" || r.URL.Path == "/sign-up" {
				http.Redirect(w, r, "/profile", http.StatusSeeOther)
				return
			}
			next(w, r)
			return
		}

		if r.URL.Path == "/sign-up" {
			templates.ExecuteTemplate(w, "sign-up.html", nil)
			return
		}

		if r.URL.Path == "/edit-profile" || r.URL.Path == "/profile" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		templates.ExecuteTemplate(w, "login.html", types.ErrorLoginType{GoogleUrl: env.FRONT_END_GOOGLE_LOGIN_URL})
	})
}
