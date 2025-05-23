package api

import (
	"encoding/json"
	// AppTokenService "go-hh/src/services/appToken"
	"net/http"
	// VacancyService "go-hh/src/services/vacancy"

	// "io"
	// "os"
)

type NegotiationsRequestData struct {
	Message string `json:"message"`
	ResumeId string `json:"resume_id"`
	VacancyIdList []string `json:"vacancy_id_list"`
}

func OtClick(w http.ResponseWriter, r *http.Request) {
	var request NegotiationsRequestData
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// ... 
}