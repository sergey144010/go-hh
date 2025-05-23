package api

import (
	"encoding/json"
	AppTokenService "go-hh/src/services/appToken"
	"net/http"
	VacancyService "go-hh/src/services/vacancy"

	// "io"
	// "os"
)

type RequestData struct {
	Code string `json:"code"`
}

type ResponseData struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

func TakeTokens(w http.ResponseWriter, r *http.Request) {
	var requestData RequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// client := &http.Client{}
	// req, _ := request()
	// resp, _ := client.Do(req)

	// io.Copy(os.Stdout, resp.Body)

	body, _ := VacancyService.RequestTake(func() (*http.Request, error) {
		return AppTokenService.RequestAccessToken(
			"https://hh.ru/oauth/token",
			AppTokenService.AUTHORIZATION_CODE,
			AppTokenService.FormDataUser(requestData.Code),
		)
	})

	var responseData ResponseData
	json.Unmarshal(body, &responseData)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(responseData)
}
