package persons

import (
	"context"
	"personService/app/projections"
)

func (qs *personQueryService) GetPersons(ctx context.Context, query GetPersonsQuery) ([]*projections.Person, error) {
	const q = `SELECT id, first_name, last_name, is_blocked FROM persons`

	var persons []*projections.Person
	err := qs.Tx.GetDB(ctx).Select(&persons, q)

	return persons, err
}

type GetPersonsQuery struct {
	Offset int64
	Limit  int32
}
