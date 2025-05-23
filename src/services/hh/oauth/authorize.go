package oauth

import (
	AppRequestService "go-hh/src/services/appRequest"
	"os"
	"fmt"
	"io"
	"net/http"
)

func Uri() string {
	clientId := os.Getenv("CLIENT_ID")
	state := "5t8wr34543fsd"
	redirectUri := os.Getenv("REDIRECT_URI")

	return "https://hh.ru/oauth/authorize?response_type=code&" +
	"client_id=" + clientId + "&" +
	"state=" + state + "&" +
	"redirect_uri=" + redirectUri
}

func Authorize() []byte {

	request, err := AppRequestService.SimpleGet(Uri())

	if err != nil {
		fmt.Println("Error request", err)
	}

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		fmt.Println("Error read body", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error read body", err)
	}

	return body
}