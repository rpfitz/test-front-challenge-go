package controller

import (
	"bytes"
	"encoding/json"
	"frontendmod/env"
	"io"
	"log"
	"net/http"
)

func GetCookie(r *http.Request, cookieName string) (string, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return cookie.Value, err
	}
	return cookie.Value, nil
}

func GETAuthUser(cookie string) (*http.Response, FetchUserResponseType, error) {
	var myResponse FetchUserResponseType
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
