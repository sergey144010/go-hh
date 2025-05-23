package AppRequestService

import (
	"net/http"
	"os"
)

func RequestGet(uri string, accessToken string) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("User-Agent", os.Getenv("USER_AGENT"))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer " + accessToken)

	return request, nil
}

func SimpleGet(uri string) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	return request, nil
}