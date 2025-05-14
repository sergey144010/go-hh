package TestController

import (
	"go-hh/src/services/baseUrlBuilder"
	"net/http"
	"fmt"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	builder := BaseUrlBuilder.BaseUrlBuilder{
		BaseUrl: "https://api.hh.ru/vacancies?",
	}

	fmt.Fprint(w, builder.Text("php").Area(113).Period(1).ResponsesCountEnabled().Get())
}
