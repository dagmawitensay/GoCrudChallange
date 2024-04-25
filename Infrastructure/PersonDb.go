package Infrastructure

import (
	"errors"
	"github.com/google/uuid"
	"github.com/dagmawitensay/GOCRUDCHALLANGE/Domain"
)

type PersonDatabase struct {
	persons []Domain.Person
}

func NewPersonDatabase(database []Domain.Person) *PersonDatabase {
	return &PersonDatabase{persons: database}
}

func (db *PersonDatabase) GetPersonById(id string)(*Domain.Person, error) {
	for _, person := range db.persons {
		if (person.Id == id) {
			return &person, nil
		}
	}

	return nil, errors.New("PERSON NOT FOUND")
}

func (db* PersonDatabase) GetAllPersons() []Domain.Person {
	return db.persons
}

func (db *PersonDatabase) CreatePerson(person Domain.Person) (*Domain.Person, error) {
	person.Id = uuid.New().String()
	db.persons = append(db.persons, person)
	return &person, nil
}

func (db *PersonDatabase) UpdatePerson(personId string, person Domain.Person) (error) {
	person.Id = personId
	
	for i, p := range db.persons {
		if p.Id == personId {
			db.persons[i] = person
			return nil
		}
	}
	return errors.New("PERSON NOT FOUND")
}

func (db *PersonDatabase) DeletePerson(id string) error {
	for i, person := range db.persons {
		if person.Id == id {
			db.persons = append(db.persons[:i], db.persons[i+1:]...)
			return nil
		}
	}

	return errors.New("PERSON NOT FOUND")
}