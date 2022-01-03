package persons

import (
	"context"
	"github.com/google/uuid"
	"personService/app/projections"
	"personService/domain"
)

func (qs *personQueryService) GetPersons(ctx context.Context, query GetPersonsQuery) ([]*projections.Person, error) {
	person := &projections.Person{PersonState: domain.PersonState{
		Id:        uuid.New(),
		FirstName: "Some",
		LastName:  "Person",
	}}
	return []*projections.Person{person}, nil
}

type GetPersonsQuery struct {
	Offset int64
	Limit  int32
}
