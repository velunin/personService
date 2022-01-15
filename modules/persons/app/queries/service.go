package queries

import (
	"context"
	"go.uber.org/fx"
	"personService/internal/database"
)

type PersonQueryService interface {
	GetPerson(ctx context.Context, query GetPersonQuery) (*Person, error)
	GetPersons(ctx context.Context, query GetPersonsQuery) ([]*Person, error)
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
