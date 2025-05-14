package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
    "github.com/gorilla/mux"
	"go-hh/src/controllers/vacancies"
	"go-hh/src/controllers/test"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")

		return
	}

	router := mux.NewRouter()
    router.HandleFunc("/vacancies", VacanciesController.Handler)
	router.HandleFunc("/test", TestController.Handler)
    http.Handle("/",router)
 
    fmt.Println("Server is listening...")
    http.ListenAndServe(":8181", nil)
}