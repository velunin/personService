package persons

import (
	"context"
	"go.uber.org/fx"
	"personService/app/database"
	"personService/app/projections"
)

type PersonQueryService interface {
	GetPerson(ctx context.Context, query GetPersonQuery) (*projections.Person, error)
	GetPersons(ctx context.Context, query GetPersonsQuery) ([]*projections.Person, error)
}

type personQueryService struct {
	QsParams
}

type QsParams struct {
	fx.In
	Tx database.Transaction
}

func NewPersonQueryService(params QsParams) PersonQueryService {
	return &personQueryService{params}
}
