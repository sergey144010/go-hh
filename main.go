package main

import (
	"fmt"
	"go-hh/src/controllers/api"
	LoginHhController "go-hh/src/controllers/login-hh"
	IndexController "go-hh/src/controllers/main"
	TestController "go-hh/src/controllers/test"
	VacanciesController "go-hh/src/controllers/vacancies"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")

		// return
	}

	router := mux.NewRouter()
	router.HandleFunc("/", IndexController.Handler)
	router.HandleFunc("/vacancies", redirectToMain)
	router.HandleFunc("/vacancies/", redirectToMain)
	router.HandleFunc("/vacancies/php/{period}", VacanciesController.Handler)
	router.HandleFunc("/test", TestController.Handler)
	router.HandleFunc("/login-hh", LoginHhController.Handler)
	router.HandleFunc("/api/take-token", api.TakeTokens).Methods(http.MethodPost)
	router.HandleFunc("/api/send", api.Send).Methods(http.MethodPost)
	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":"+port(), nil)
}

func port() string {
	return os.Getenv("WEB_SERVER_PORT")
}

func redirectToMain(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
