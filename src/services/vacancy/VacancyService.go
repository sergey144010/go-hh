package VacancyService

import (
	"encoding/json"
	"fmt"
	response "go-hh/src/responses"
	BaseUrlBuilder "go-hh/src/services/baseUrlBuilder"
	"io"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
)

type TakeResponse struct {
	ListVacancies  *[]response.Vacancy
	Count          int
	WorkFormatList *map[string]string
}

func Take() TakeResponse {
	var listVacancies []response.Vacancy
	vacanciesRawFirst, _ := requestTake(func() (*http.Request, error) {
		return requestGetVacancies(0)
	})
	vacanciesFirst := vacanciesCollection(vacanciesRawFirst)
	listVacancies = append(listVacancies, vacanciesFirst.Items...)

	for i := 1; i < vacanciesFirst.PerPage; i++ {
		vacanciesRaw, _ := requestTake(func() (*http.Request, error) {
			return requestGetVacancies(i)
		})
		vacancies := vacanciesCollection(vacanciesRaw)
		listVacancies = append(listVacancies, vacancies.Items...)
	}

	sort.SliceStable(listVacancies, func(i, j int) bool {
		return listVacancies[i].Counters.TotalResponses < listVacancies[j].Counters.TotalResponses
	})

	workFormatList := map[string]string{}
	for _, vacancy := range listVacancies {
		for _, item := range vacancy.WorkFormat {
			if _, ok := workFormatList[item.Id]; !ok {
				workFormatList[item.Id] = item.Name
			}
		}
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

func requestGet(uri string) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("User-Agent", os.Getenv("USER_AGENT"))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+accessToken())

	return request, nil
}

func baseUrl() string {
	// return "https://api.hh.ru/vacancies?text=php&area=113&period=1&responses_count_enabled=true"

	builder := BaseUrlBuilder.BaseUrlBuilder{
		BaseUrl: "https://api.hh.ru/vacancies?",
	}

	return builder.Text("php").Area(113).Period(1).ResponsesCountEnabled().Get()
}

func requestGetVacancies(page int) (*http.Request, error) {
	return requestGet(baseUrl() + "&page=" + fmt.Sprintf("%d", page))
}
func requestGetAreas() (*http.Request, error) {
	return requestGet("https://api.hh.ru/areas")
}

func requestTake(calback func() (*http.Request, error)) ([]byte, error) {
	client := &http.Client{}

	request, err := calback()

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

func areasCollection(body []byte) {
	var areasCollection []response.AreaRoot
	errUn := json.Unmarshal(body, &areasCollection)
	if errUn != nil {
		fmt.Println(errUn)
	}

	fmt.Println(areasCollection)
}

func formData() url.Values {
	formData := url.Values{}
	formData.Set("grant_type", "client_credentials")
	formData.Set("client_id", os.Getenv("CLIENT_ID"))
	formData.Set("client_secret", os.Getenv("CLIENT_SECRET"))

	return formData
}

func requestAccessToken() (*http.Request, error) {
	formData := formData()
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
