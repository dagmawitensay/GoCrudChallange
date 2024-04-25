package Application

import (
	"errors"
	"github.com/dagmawitensay/GOCRUDCHALLANGE/Domain"
	"github.com/dagmawitensay/GOCRUDCHALLANGE/Infrastructure"
)

type PersonService struct {
	personDatabase *Infrastructure.PersonDatabase
}

func NewPersonService(db *Infrastructure.PersonDatabase) *PersonService {
	return &PersonService{personDatabase: db}
}

func (ps *PersonService) GetPersonById(id string) (*Domain.Person, error) {
	person, err := ps.personDatabase.GetPersonById(id)
	if person == nil {
		return nil, errors.New(err.Error())
	}
	return person, nil
}

func (ps *PersonService) GetAllPersons() []Domain.Person {
	return ps.personDatabase.GetAllPersons()
}

func (ps *PersonService) CreatePerson(person Domain.Person) (*Domain.Person, error) {
	return ps.personDatabase.CreatePerson(person)
}

func (ps *PersonService) UpdatePerson(personId string, person Domain.Person) error {
	err := ps.personDatabase.UpdatePerson(personId, person)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PersonService) DeletePerson(id string) error {
	err := ps.personDatabase.DeletePerson(id)
	if err != nil {
		return err
	}

	return nil
}