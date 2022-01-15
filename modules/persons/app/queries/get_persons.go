package queries

import (
	"context"
)

func (qs *personQueryService) GetPersons(ctx context.Context, query GetPersonsQuery) ([]*Person, error) {
	const q = `SELECT * FROM persons LIMIT $1 OFFSET $2`

	limit := int32(10)
	if query.Limit > 0 {
		limit = query.Limit
	}

	var persons []*Person
	err := qs.Tx.GetDB(ctx).Select(&persons, q, limit, query.Offset)

	return persons, err
}

type GetPersonsQuery struct {
	Offset int64
	Limit  int32
}
