package controller

import (
	"net/http"
)

func GetEditProfile(w http.ResponseWriter, r *http.Request) {
	_, email, full_name, telephone, isGoogle := GetCookiesSession(r)

	templates.ExecuteTemplate(w, "edit-profile.html", UpdateProfileSuccessType{
		Email:     email,
		Full_Name: full_name,
		Telephone: telephone,
		Google:    isGoogle,
	})
	return
}

func PostEditProfile(w http.ResponseWriter, r *http.Request) {
	_, profileEdited, _, emailExists := HandleEditProfile(w, r)

	if profileEdited {
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return
	}

	if emailExists {
		_, email, full_name, telephone, google := GetCookiesSession(r)
		templates.ExecuteTemplate(w, "edit-profile.html", UpdateProfileSuccessType{
			Full_Name: full_name,
			Telephone: telephone,
			Email:     email,
			Google:    google,
			Error:     "Email already exists.",
		})
		return
	}

	http.Redirect(w, r, "/profile", http.StatusSeeOther)
	return
}
