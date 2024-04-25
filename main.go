package main

import (
	"github.com/dagmawitensay/GOCRUDCHALLANGE/Application"
	"github.com/dagmawitensay/GOCRUDCHALLANGE/Domain"
	"github.com/dagmawitensay/GOCRUDCHALLANGE/Infrastructure"
	"github.com/dagmawitensay/GOCRUDCHALLANGE/WebApi"
	"github.com/gorilla/mux"
	"net/http"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()
	router.Use(enableCORS)
	personDatabase := Infrastructure.NewPersonDatabase([]Domain.Person{})
	personService := Application.NewPersonService(personDatabase)
	personController := WebApi.NewPersonController(personService)
	personController.SetupRotues(router)

	router.NotFoundHandler = http.HandlerFunc(personController.NotFoundHandler)
	http.Handle("/", router)

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		print(err)
	}
}