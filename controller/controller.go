package controller

import (
	"bytes"
	"encoding/json"
	"frontendmod/env"
	"html/template"
	"io"
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

func FetchPostAPILogin(email string, password string) (string, AuthenticatedUserType) {
	method := "POST"
	httpposturl := env.BACK_END_POST_LOGIN_URL

	payload := map[string]interface{}{"email": email, "password": password}
	byts, _ := json.Marshal(payload)

	request, error := http.NewRequest(method, httpposturl, bytes.NewBuffer(byts))
	if error != nil {
		log.Println("FetchPostAPILogin - Error making request: ", error)
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		log.Println("FetchPostAPILogin - Error executing request: ", error)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	var myResponse AuthenticatedUserType
	json.Unmarshal(body, &myResponse)
	return response.Status, myResponse
}

func FetchPostAPISignUp(email string, password string) (string, AuthenticatedUserType) {
	method := "POST"
	httpposturl := env.BACK_END_POST_SIGN_UP_URL

	payload := map[string]interface{}{"email": email, "password": password}
	byts, _ := json.Marshal(payload)

	request, error := http.NewRequest(method, httpposturl, bytes.NewBuffer(byts))
	if error != nil {
		log.Println("FetchPostAPISignUp - Error making request: ", error)
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		log.Println("FetchPostAPISignUp - Error executing request: ", error)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	var myResponse AuthenticatedUserType
	json.Unmarshal(body, &myResponse)
	return response.Status, myResponse
}
