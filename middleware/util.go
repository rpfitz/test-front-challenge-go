package middleware

import (
	"encoding/json"
	"frontendmod/env"
	"io"
	"log"
	"net/http"
)

type AuthUserResponseType struct {
	Success bool         `json:"id,omitempty"`
	User    AuthUserType `json:"user,omitempty"`
}

type AuthUserType struct {
	ID        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Full_Name string `json:"full_name,omitempty"`
	Telephone string `json:"telephone,omitempty"`
	Password  string `json:"password,omitempty"`
	Google    string `json:"google,omitempty"`
}

type ErrorLoginType struct {
	Error     string
	GoogleUrl string
}

func IsUserAuthenticated(w http.ResponseWriter, r *http.Request) (AuthUserResponseType, bool) {
	var failedResponse AuthUserResponseType

	authorizationCookie, err := GetCookie(r, "Authorization")
	if err != nil {
		return failedResponse, false
	}

	_, response, err := GETAuthUser(authorizationCookie)

	if response.Success == true {
		return response, true
	}

	return failedResponse, false
}

type GETAuthUserType struct {
	ID   string              `json:"id,omitempty"`
	User GETAuthUserTypeUser `json:"user,omitempty"`
}

type GETAuthUserTypeUser struct {
	Id        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Full_Name string `json:"full_name,omitempty"`
	Telephone string `json:"telephone,omitempty"`
	Password  string `json:"password,omitempty"`
	Google    string `json:"google,omitempty"`
}

func GETAuthUser(authorizationCookie string) (*http.Response, AuthUserResponseType, error) {
	var myResponse AuthUserResponseType
	url := env.BACK_END_GET_AUTH_USER_URL

	req, err := http.NewRequest("GET", url, nil)
	req.AddCookie(&http.Cookie{
		Name:  "Authorization",
		Value: authorizationCookie,
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("GETAuthUser - Error making request:", err)
		return resp, myResponse, nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &myResponse)
	return resp, myResponse, nil
}

func SetMyCacheHeaders(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}

func GetCookie(r *http.Request, cookieName string) (string, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}
