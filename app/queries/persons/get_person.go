package persons

import (
	"context"
	"github.com/google/uuid"
	"personService/app/projections"
	"personService/domain"
)

func (qs *personQueryService) GetPerson(ctx context.Context, query GetPersonQuery) (*projections.Person, error) {
	return &projections.Person{
		PersonState: domain.PersonState{
			Id:        uuid.New(),
			FirstName: "Some",
			LastName:  "Person",
		},
	}, nil
}

type GetPersonQuery struct {
	Id uuid.UUID
}
