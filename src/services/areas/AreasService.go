package AreasService

import (
	"net/http"
	AppRequestService "go-hh/src/services/appRequest"
)

func RequestGetAreas(accessToken string) (*http.Request, error) {
	return AppRequestService.RequestGet(
		"https://api.hh.ru/areas",
		accessToken,
	)
}