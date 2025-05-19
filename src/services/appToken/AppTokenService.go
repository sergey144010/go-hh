package appTokenService

import (
	"encoding/json"
	"fmt"
	response "go-hh/src/responses"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func AreasCollection(body []byte) {
	var areasCollection []response.AreaRoot
	errUn := json.Unmarshal(body, &areasCollection)
	if errUn != nil {
		fmt.Println(errUn)
	}

	fmt.Println(areasCollection)
}

func FormData() url.Values {
	formData := url.Values{}
	formData.Set("grant_type", "client_credentials")
	formData.Set("client_id", os.Getenv("CLIENT_ID"))
	formData.Set("client_secret", os.Getenv("CLIENT_SECRET"))

	return formData
}

func RequestAccessToken() (*http.Request, error) {
	formData := FormData()
	reqBody := strings.NewReader(formData.Encode())
	uri := "https://api.hh.ru/token"

	request, err := http.NewRequest(http.MethodPost, uri, reqBody)
	if err != nil {
		return nil, err
	}
	request.Header.Add("User-Agent", os.Getenv("USER_AGENT"))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return request, nil
}