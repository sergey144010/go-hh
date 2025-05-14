package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
    "github.com/gorilla/mux"
	"go-hh/src/controllers/vacancies"
	"go-hh/src/controllers/test"
	"os"
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
    http.ListenAndServe(":" + port(), nil)
}

func port() string {
	return os.Getenv("WEB_SERVER_PORT")
}