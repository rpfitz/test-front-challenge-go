package service

import (
	"bytes"
	"encoding/json"
	"frontendmod/env"
	"frontendmod/types"
	"io"
	"log"
	"net/http"
)

func FetchAPI(method string, url string, payload map[string]interface{}) ([]byte, error) {
	byts, _ := json.Marshal(payload)

	request, error := http.NewRequest(method, url, bytes.NewBuffer(byts))
	if error != nil {
		log.Println("fetchAPI - Error making request.")
		return []byte("Error making request."), error
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}

	response, error := client.Do(request)
	if error != nil {
		log.Println("fetchAPI - Error executing request.")
		return []byte("Error executing request."), error
	}

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	if error != nil {
		log.Println("fetchAPI - Error ReadAll.")
		return body, error
	}

	return body, nil
}

func FetchPOSTEditProfileAPI(full_name, telephone, email, userCookieToken string) (string, types.UpdateProfileSuccessType) {
	var responseBody types.UpdateProfileSuccessType
	payload := map[string]interface{}{"full_name": full_name, "telephone": telephone, "email": email}
	url := env.BACK_END_POST_EDIT_PROFILE_URL

	byts, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(byts))

	req.AddCookie(&http.Cookie{
		Name:  "Authorization",
		Value: userCookieToken,
	})

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		return response.Status, types.UpdateProfileSuccessType{}
	}

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	json.Unmarshal(body, &responseBody)
	return response.Status, responseBody
}

func GETAuthUser(cookie string) (*http.Response, types.FetchUserResponseType, error) {
	var myResponse types.FetchUserResponseType
	url := env.BACK_END_GET_AUTH_USER_URL

	req, err := http.NewRequest("GET", url, nil)
	req.AddCookie(&http.Cookie{
		Name:  "Authorization",
		Value: cookie,
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return resp, myResponse, nil
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	json.Unmarshal(body, &myResponse)
	return resp, myResponse, nil
}

func GETAuthUser2(authorizationCookie string) (*http.Response, types.AuthUserResponseType, error) {
	var myResponse types.AuthUserResponseType
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

func BytePOSTRequest(payload []byte, path string, endpoint string) (int, []byte) {
	url := path + endpoint

	request, error := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if error != nil {
		log.Println("MakePostRequest - Error making request: ", error)
	}

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		log.Println("MakePostRequest - Error executing request: ", error.Error())
		return response.StatusCode, nil
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	return response.StatusCode, body
}

func MakePostRequest(
	req *http.Request,
	payload map[string]interface{},
	path string,
	endpoint string,
) (int, []byte) {
	// payloadExample := map[string]interface{}{"full_name": full_name, "telephone": telephone, "email": email}
	// enpointExample := "/edit-profile?email="
	// queryParamExample := "value"
	// httpposturlExammple := "http://" + env.host + ":" + env.port + endpoint
	// json.Unmarshal(body, &responseBody)

	url := path + endpoint

	byts, _ := json.Marshal(payload)

	request, error := http.NewRequest("POST", url, bytes.NewBuffer(byts))
	if error != nil {
		log.Println("MakePostRequest - Error making request: ", error)
	}

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		log.Println("MakePostRequest - Error executing request: ", error.Error())
		return response.StatusCode, nil
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	return response.StatusCode, body
}
