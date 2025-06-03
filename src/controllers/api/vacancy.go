package api

import (
	"encoding/json"
	"net/http"
	VacancyService "go-hh/src/services/vacancy"
	"fmt"
	"bytes"
	"io"
	"os"
	"strings" 
)

type AnalyzeRequestData struct {
	VacancyIdList []string `json:"vacancy_id_list"`
	AccessToken string `json:"access_token"`
}

type VacancyResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

type VacancyContainList struct {
	Id string `json:"id"`
	Cover bool `json:"cover"`
}

func Analyze(w http.ResponseWriter, r *http.Request) {
	var request AnalyzeRequestData
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var vacancyList []VacancyResponse
	for _, vacancyId := range request.VacancyIdList {
		body, _ := VacancyService.RequestTake(func() (*http.Request, error) {
			return VacancyService.RequestGetVacancy(vacancyId, request.AccessToken)
		})
		var vacancy VacancyResponse
		errUn := json.Unmarshal(body, &vacancy)
		if errUn != nil {
			fmt.Println(errUn)
		}
	
		vacancyList = append(vacancyList, vacancy)
	}

	var vacancyCoverList []VacancyContainList
	for _, vacancy := range vacancyList {
		if (strings.Contains(vacancy.Description, "сопрово")) {
			vacancyCoverList = append(vacancyCoverList,	VacancyContainList{Id: vacancy.Id,Cover: true})

			continue
		}

		vacancyCoverList = append(vacancyCoverList,	VacancyContainList{Id: vacancy.Id,Cover: false})
	}

	// data, _ := json.Marshal(vacancyList)
	// reader := bytes.NewReader(data)
	// io.Copy(os.Stdout, reader)

	data1, _ := json.Marshal(vacancyCoverList)
	reader1 := bytes.NewReader(data1)
	io.Copy(os.Stdout, reader1)
}