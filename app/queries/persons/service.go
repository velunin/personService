package persons

import (
	"context"
	"personService/app/projections"
)

type PersonQueryService interface {
	GetPerson(ctx context.Context, query GetPersonQuery) (*projections.Person, error)
	GetPersons(ctx context.Context, query GetPersonsQuery) ([]*projections.Person, error)
}

type personQueryService struct {
}

func NewPersonQueryService() PersonQueryService {
	return &personQueryService{}
}
