package VacancyService

import (
	"encoding/json"
	"fmt"
	response "go-hh/src/responses"
	AppRequestService "go-hh/src/services/appRequest"
	BaseUrlBuilder "go-hh/src/services/baseUrlBuilder"
	"io"
	"net/http"
	"os"
	"sort"
	"time"
)

type TakeResponse struct {
	ListVacancies  *[]response.Vacancy
	Count          int
	WorkFormatList *map[string]string
}

func Take(period string) TakeResponse {
	var listVacancies []response.Vacancy
	vacanciesRawFirst, _ := RequestTake(func() (*http.Request, error) {
		return requestGetVacancies(0, period)
	})
	vacanciesFirst := vacanciesCollection(vacanciesRawFirst)
	listVacancies = append(listVacancies, vacanciesFirst.Items...)

	for i := 1; i < vacanciesFirst.PerPage; i++ {
		vacanciesRaw, _ := RequestTake(func() (*http.Request, error) {
			return requestGetVacancies(i, period)
		})
		vacancies := vacanciesCollection(vacanciesRaw)
		listVacancies = append(listVacancies, vacancies.Items...)
	}

	sort.SliceStable(listVacancies, func(i, j int) bool {
		return listVacancies[i].Counters.TotalResponses < listVacancies[j].Counters.TotalResponses
	})

	workFormatList := map[string]string{}
	for index, vacancy := range listVacancies {
		for _, item := range vacancy.WorkFormat {
			if _, ok := workFormatList[item.Id]; !ok {
				workFormatList[item.Id] = item.Name
			}
		}
		t, _ := time.Parse("2006-01-02T15:04:05+0300", vacancy.CreatedAt)
		listVacancies[index].CreatedAt = t.Format(time.DateOnly)
	}

	return TakeResponse{
		ListVacancies:  &listVacancies,
		Count:          vacanciesFirst.Found,
		WorkFormatList: &workFormatList,
	}
}

func accessToken() string {
	return os.Getenv("ACCESS_TOKEN")
}

func baseUrl(periodInput string) string {
	// return "https://api.hh.ru/vacancies?text=php&area=113&period=1&responses_count_enabled=true"

	builder := BaseUrlBuilder.BaseUrlBuilder{
		BaseUrl: "https://api.hh.ru/vacancies?",
	}

	period := 1
	if periodInput == "day" {
		period = 1
	}
	if periodInput == "month" {
		period = 30
	}

	return builder.Text("php").Area(113).Period(period).ResponsesCountEnabled().Get()
}

func requestGetVacancies(page int, period string) (*http.Request, error) {
	return AppRequestService.RequestGet(
		baseUrl(period)+"&page="+fmt.Sprintf("%d", page),
		accessToken(),
	)
}

func RequestTake(callback func() (*http.Request, error)) ([]byte, error) {
	client := &http.Client{}

	request, err := callback()

	if err != nil {
		fmt.Println("Error request", err)

		return nil, err
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("Error read body", err)

		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error read body", err)

		return nil, err
	}

	return body, nil
}

func vacanciesCollection(body []byte) *response.VacanciesResponse {
	var vacancies response.VacanciesResponse
	errUn := json.Unmarshal(body, &vacancies)
	if errUn != nil {
		fmt.Println(errUn)
	}

	return &vacancies
}
