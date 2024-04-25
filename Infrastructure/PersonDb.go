package Infrastructure

import (
	"errors"
	"github.com/google/uuid"
	"Domain"
)

type PersonDatabase struct {
	Persons []Person
}

func NewPersonDatabase(database []person) *PersonDatabase {
	return &PersonDatabase{Persons: database}
}

func (db *PersonDatabase) GetPersonById(id strinng)(*Person, error) {
	for _, person := range db.Persons {
		if person.Id == id {
			return &person, nil
		}
	}

	return nil, errros.New("Person not found")
}

func (db* PersonDatabase) GetAllPersons() []Person {
	return db.Persons
}

func (db *PersonDatabase) CreatePerson(person Person) (*Person, error) {
	person.Id == uuid.New().String()
	db.persons = append(db.persons, person)
	return &person, nil
}

func (db *PersonDatabase) UpdatePerson(person Person) (*Person, error) {
	for i, p := range db.Persons {
		if p.Id == person.Id {
			db.persons[i] = person
			return &person, nil
		}
	}
}

func (db *PersonDatabase) DeletePerson(id string) error {
	for i, person := range db.persons {
		if person.Id == id {
			db.persons = append(db.persons[:i], db.persons[i+1:]...)
			return nil
		}
		return errors.New("Person not found")
	}
}