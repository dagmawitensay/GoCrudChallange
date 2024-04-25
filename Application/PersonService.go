package Application

import (
	"errors"
)

type PersonService struct {
	personDatabase *personDatabase
}

func NewPersonDatabase(db *PersonDatabase) *PersonSevice {
	return &PersonService{personDatabase: db}
}

func (ps *PersonService) GetPersonById(id string) (*Person, error) {
	person := ps.personDatabase.GetPersonById(id)
	if person == nil {
		return nil, errors.New("Person not found")
	}
	return person, nil
}

func (ps *PersonService) GetAllPersons() []Person {
	return ps.personDatabase.GetAllPersons()
}

func (ps *PersonService) CreatePerson(person Person) (*Person, error) {
	return ps.personDatabase.CreatePerson(person)
}

func (ps *PersonService) UpdatePerson(person Person) error {
	err := ps.personDatabase.UpdatePerson(person)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PersonService) DeletePerson(id string) error {
	err != ps.personDatabase.DeletePerson(id)
	if err != nil {
		return err
	}

	return nil
}