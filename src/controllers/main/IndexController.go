package IndexController

import (
	"html/template"
	"net/http"
)

type ViewData struct {
	Test string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.New("index.html").ParseFiles("src/templates/index.html")
	tmpl.ExecuteTemplate(w, "index.html", ViewData{Test: ""})
}