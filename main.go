package main

import (
	"github.com/dagmawitensay/Application"
	"github.com/dagmawitensay/Domain"
	"github.com/dagmawitensay/Infrastructure"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRotuer()
	personDatabase := NewPersonDatabase()
	personService := NewPersonService(personDatabase)
	personController := NewPersonController(personService)
	personController.SetUpRoute(routes)

	router.NotFoundHandler = http.HandlerFunc(personController.NotFoundHandler)
	http.Handle("/", router)

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		print(err)
	}
}