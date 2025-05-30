package VacanciesController

import (
	res "go-hh/src/responses"
	"go-hh/src/services/hh/oauth"
	VacancyService "go-hh/src/services/vacancy"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type ViewData struct {
	Vacancies *[]res.Vacancy
	Count int
	WorkFormatList *[]res.IdName
	Uri string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    period := vars["period"]
	response := VacancyService.Take(period)

	workFormatList := []res.IdName{}

	for key, value := range *response.WorkFormatList {
		workFormatList = append(workFormatList, res.IdName{Id: key, Name: value})
	}

	data := ViewData{
		Vacancies: response.ListVacancies,
		Count: response.Count,
		WorkFormatList: &workFormatList,
		Uri: oauth.Uri(),
	}

	// var vacancies []response.Vacancy
	// vacancies = append(vacancies, response.Vacancy{
	// 	Name: "one",
	// })
	// vacancies = append(vacancies, response.Vacancy{
	// 	Name: "two",
	// })
	// vacancies = append(vacancies, response.Vacancy{
	// 	Name: "three",
	// })
	// vacancies = append(vacancies, response.Vacancy{
	// 	Name: "one two",
	// })
	// vacancies = append(vacancies, response.Vacancy{
	// 	Name: "one two three",
	// })
	// vacancies = append(vacancies, response.Vacancy{
	// 	Name: "two three",
	// })
	// data := ViewData{
	// 	Vacancies: &vacancies,
	// }

	tmpl, _ := template.New("vacancies.html").Delims("[[", "]]").ParseFiles("src/templates/vacancies.html")
	tmpl.ExecuteTemplate(w, "vacancies.html", &data)
}
