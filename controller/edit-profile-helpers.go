package controller

import (
	"bytes"
	"encoding/json"
	"frontendmod/env"
	"io"
	"log"
	"net/http"
)

type UpdateProfileSuccessType struct {
	Success            bool   `json:"success"`
	Message            string `json:"message"`
	Email_From_Request string `json:"email_from_request"`
	Full_Name          string `json:"full_name"`
	Telephone          string `json:"telephone"`
	Email              string `json:"email"`
	Full_Name_Updated  bool   `json:"full_name_updated"`
	Telephone_Updated  bool   `json:"telephone_updated"`
	Email_Updated      bool   `json:"email_updated"`
	Token              string `json:"token"`
	Google             string `json:"google"`
	Error              string `json:"error"`
	Done               string `json:"done"`
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

func whatToEditProfile(userCookieEmail string, email string, userCookieFullName string, full_name string, userCookieTelephone string, telephone string) (bool, bool, bool) {
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

func HandleEditProfile(w http.ResponseWriter, r *http.Request) (UpdateProfileSuccessType, bool, string, bool) {
	var failedResponse UpdateProfileSuccessType

	full_name := r.FormValue("full_name")
	telephone := r.FormValue("telephone")
	email := r.FormValue("email")

	userCookieToken, userCookieEmail, userCookieFullName, userCookieTelephone, _ := GetCookiesSession((r))

	editEmail, editFullName, editTelephone := whatToEditProfile(userCookieEmail, email, userCookieFullName, full_name, userCookieTelephone, telephone)
	if (editFullName == false) && (editEmail == false) && (editTelephone == false) {
		return failedResponse, false, "??", false
	}

	status, response := fetchPOSTEditProfileAPI(full_name, telephone, email, userCookieToken)
	if status != "200 OK" {
		return response, false, "Server error.", false
	}

	if response.Success {
		cookiesUpdateEditProfile(w, response, email, full_name, telephone, response.Token)
		return response, true, response.Message, false
	}

	if !response.Success && response.Message == "Email already exists." {
		return response, false, response.Message, true
	}

	return failedResponse, false, "Something went wrong.", false
}

func fetchPOSTEditProfileAPI(full_name, telephone, email, userCookieToken string) (string, UpdateProfileSuccessType) {
	var responseBody UpdateProfileSuccessType

	payload := map[string]interface{}{"full_name": full_name, "telephone": telephone, "email": email}
	byts, _ := json.Marshal(payload)

	url := env.BACK_END_POST_EDIT_PROFILE_URL

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(byts))
	req.AddCookie(&http.Cookie{
		Name:  "Authorization",
		Value: userCookieToken,
	})

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return response.Status, UpdateProfileSuccessType{}
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	json.Unmarshal(body, &responseBody)
	return response.Status, responseBody
}

func cookiesUpdateEditProfile(w http.ResponseWriter, responseBody UpdateProfileSuccessType, email string, full_name string, telephone string, newToken string) {
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
