package WebApi

import (
	"encoind/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

type PersonController struct {
	personService *PersonService
}

func NewPersonController(personService *PersonService) *PersonController {
	return &PersonController{personService: personService}
}

func (pc *PersonController) SetupRotues(router *mux.Router) {
	router.HandleFunc("/person", pc.GetAllPersons).Methods("GET")
	router.HandleFunc("/person/{personId}", pc.GetPersonById).Methods("GET")
	router.HandleFunc("/person", pc.CreatePerson).Methods("POST")
	router.HandleFunc("/person/{personId}", pc.UpdatePerson).Methods("PUT")
	router.HandleFunc("/person/{personId}", pc.DeltePerson).Methods("DELETE")
}

func (pc *PersonController) GetAllPersons(w http.ResponseWriter, r *http.Request) {
	allPersons := pc.personService.GetAllPersons()
	json.NewEncoder(w).Encode(allPersons)
}

func (pc *PersonController) GetPersonById(w http.ResponseWriter, r *Http.Request) {
	params := mux.Vars(r)
	personId := params["personId"]
	person, err != pc.personService.GetPersonById(personId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error" : err.Error()})
		return
	}
	json.NewEncoder(w).Encode(person)
}

func (pc *PersonController) CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	createdPerson, err := pc.personService.CreatePerson(person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createdPerson)

}

func (pc *PersonController) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personId := params["personId"]
	var person Person
	err != json.NewDecoder(r.Body).Decode(&person)
	if (err != nil) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	err := pc.personService.UpdatePerson(person)
	
	if (err != nil) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (pc *PersonController) DeltePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personId := params["personId"]

	err := pc.personService.DeletePerson(personId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Erro()})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (pc *PersonController) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Path " + r.URL.Path + " does not exist for " + r.Method + " method"})
}

func (pc *PersonController) ErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	statusCode := http.StatusInternalServerError
	if e, ok := err.(*AppError); ok {
		statusCode = e.StatusCode
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}
