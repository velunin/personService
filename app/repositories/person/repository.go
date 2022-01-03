package person

import (
	"fmt"
	"github.com/google/uuid"
	"personService/domain"
)

type Repository interface {
	Get(id uuid.UUID) (*domain.Person, error)
	List() ([]*domain.Person, error)
	Insert(person *domain.Person) error
	Update(person *domain.Person) error
	Delete(id uuid.UUID) error
}

type personRepository struct {
}

func (r *personRepository) Insert(person *domain.Person) error {
	panic("implement me")
}

func (r *personRepository) Update(person *domain.Person) error {
	panic("implement me")
}

func (r *personRepository) Delete(id uuid.UUID) error {
	panic("implement me")
}

func (r *personRepository) List() ([]*domain.Person, error) {
	return []*domain.Person{
		domain.NewPerson(uuid.New(), "3dsf", "sdfdf"),
	}, nil
}

func (r *personRepository) Get(id uuid.UUID) (*domain.Person, error) {
	return domain.NewPerson(id, "firstName", "lastName"), nil
}

func NewPersonRepository() Repository {
	fmt.Println("Create repo")
	return &personRepository{}
}
