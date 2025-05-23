package AppTokenService

import (
	"encoding/json"
	"fmt"
	response "go-hh/src/responses"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const CLIENT_CREDENTIALS = "client_credentials"
const AUTHORIZATION_CODE = "authorization_code"

func AreasCollection(body []byte) {
	var areasCollection []response.AreaRoot
	errUn := json.Unmarshal(body, &areasCollection)
	if errUn != nil {
		fmt.Println(errUn)
	}

	fmt.Println(areasCollection)
}

func FormData(grantType string) url.Values {
	formData := url.Values{}
	formData.Set("grant_type", grantType)
	formData.Set("client_id", os.Getenv("CLIENT_ID"))
	formData.Set("client_secret", os.Getenv("CLIENT_SECRET"))

	return formData
}

func FormDataApp() url.Values {
	return FormData(CLIENT_CREDENTIALS)
}

func FormDataUser(code string) url.Values {
	formData := FormData(AUTHORIZATION_CODE)
	formData.Set("code", code)
	formData.Set("redirect_uri", os.Getenv("REDIRECT_URI"))

	return formData
}

func ApiAppTokenUri() string {
	return os.Getenv("API_URI") + "/token"
}

func ApiUserTokenUri() string {
	return os.Getenv("API_URI") + "/oauth/token"
}

func RequestAccessToken(uri string, grantType string, formData url.Values) (*http.Request, error) {
	reqBody := strings.NewReader(formData.Encode())

	request, err := http.NewRequest(http.MethodPost, uri, reqBody)
	if err != nil {
		return nil, err
	}
	request.Header.Add("User-Agent", os.Getenv("USER_AGENT"))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if (grantType == AUTHORIZATION_CODE) {
		request.Header.Add("Authorization", "Bearer " + os.Getenv("ACCESS_TOKEN"))
	}

	return request, nil
}