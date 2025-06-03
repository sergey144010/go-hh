package AreasService

import (
	"net/http"
	AppRequestService "go-hh/src/services/appRequest"
	"os"
)

func RequestGetAreas(accessToken string) (*http.Request, error) {
	return AppRequestService.RequestGet(
		os.Getenv("API_URI") + "/areas",
		accessToken,
	)
}