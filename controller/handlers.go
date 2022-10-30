package controller

import (
	"encoding/json"
	"frontendmod/env"
	"frontendmod/service"
	"frontendmod/types"
	"log"
	"net/http"
)

func HandlePostLoginRequest(w http.ResponseWriter, r *http.Request) (types.SuccessLoginResponseType, error) {
	var response types.SuccessLoginResponseType

	payload := map[string]interface{}{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	}

	body, err := service.FetchAPI("POST", env.BACK_END_POST_LOGIN_URL, payload)
	if err != nil {
		log.Println("Failed to fetch data from API.")
		return response, err
	}

	err1 := json.Unmarshal(body, &response)
	if err1 != nil {
		log.Println("Failed to unmarshal type.")
		return response, err1
	}

	if response.Success {
		SetAllCookies(w, response)
		return response, nil
	}

	return response, nil
}

func HandlePostSignUpRequest(w http.ResponseWriter, r *http.Request) (types.AuthenticatedUserType, error) {
	var response types.AuthenticatedUserType

	payload := map[string]interface{}{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	}

	body, err := service.FetchAPI("POST", env.BACK_END_POST_SIGN_UP_URL, payload)
	if err != nil {
		log.Println("Failed to fetch data from API.")
		return response, err
	}

	err1 := json.Unmarshal(body, &response)
	if err1 != nil {
		log.Println("Failed to unmarshal type.")
		return response, err1
	}

	if response.Success {
		SetAllCookiesSignUp(w, response)
		return response, nil
	}

	return response, nil
}

func HandleEditProfile(w http.ResponseWriter, r *http.Request) (types.UpdateProfileSuccessType, bool, string, bool) {
	var failedResponse types.UpdateProfileSuccessType

	full_name := r.FormValue("full_name")
	telephone := r.FormValue("telephone")
	email := r.FormValue("email")

	userCookieToken, userCookieEmail, userCookieFullName, userCookieTelephone, _ := GetCookiesSession((r))

	editEmail, editFullName, editTelephone := WhatToEditProfile(userCookieEmail, email, userCookieFullName, full_name, userCookieTelephone, telephone)
	if (editFullName == false) && (editEmail == false) && (editTelephone == false) {
		return failedResponse, false, "??", false
	}

	status, response := service.FetchPOSTEditProfileAPI(full_name, telephone, email, userCookieToken)
	if status != "200 OK" {
		return response, false, "Server error.", false
	}

	if response.Success {
		CookiesUpdateEditProfile(w, response, email, full_name, telephone, response.Token)
		return response, true, response.Message, false
	}

	if !response.Success && response.Message == "Email already exists." {
		return response, false, response.Message, true
	}

	return failedResponse, false, "Something went wrong.", false
}

func HandleGoogleLogin(w http.ResponseWriter, callBackData []byte) (int, types.SignUpSuccesfulResponseType) {
	var response types.SignUpSuccesfulResponseType

	responseStatusCode, res := service.BytePOSTRequest(callBackData, env.BACK_END_HOST_URL, env.BACK_END_GOOGLE_AUTH_ENDPOINT)

	json.Unmarshal(res, &response)

	if response.Success {
		SetAllGoogleCookies(w, response)
	}

	return responseStatusCode, response
}
