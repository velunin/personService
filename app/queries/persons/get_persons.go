package persons

import (
	"context"
	"personService/app/projections"
)

func (qs *personQueryService) GetPersons(ctx context.Context, query GetPersonsQuery) ([]*projections.Person, error) {
	const q = `SELECT * FROM persons LIMIT $1 OFFSET $2`

	limit := int32(10)
	if query.Limit > 0 {
		limit = query.Limit
	}

	var persons []*projections.Person
	err := qs.Tx.GetDB(ctx).Select(&persons, q, limit, query.Offset)

	return persons, err
}

type GetPersonsQuery struct {
	Offset int64
	Limit  int32
}
