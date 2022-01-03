package person

import (
	"context"
	"github.com/google/uuid"
	"personService/domain"
)

type Repository interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.Person, error)
	List(ctx context.Context) ([]*domain.Person, error)
	Insert(ctx context.Context, person *domain.Person) error
	Update(ctx context.Context, person *domain.Person) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type personRepository struct {
}

func (r *personRepository) Insert(ctx context.Context, person *domain.Person) error {
	panic("implement me")
}

func (r *personRepository) Update(ctx context.Context, person *domain.Person) error {
	panic("implement me")
}

func (r *personRepository) Delete(ctx context.Context, id uuid.UUID) error {
	panic("implement me")
}

func (r *personRepository) List(ctx context.Context) ([]*domain.Person, error) {
	return []*domain.Person{
		domain.NewPerson(uuid.New(), "3dsf", "sdfdf"),
	}, nil
	panic("not implemented")
}

func (r *personRepository) Get(ctx context.Context, id uuid.UUID) (*domain.Person, error) {
	return domain.NewPerson(id, "firstName", "lastName"), nil
}

func NewPersonRepository() Repository {
	return &personRepository{}
}
