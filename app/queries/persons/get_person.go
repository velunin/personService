package persons

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"personService/app/projections"
	repoperson "personService/app/repositories/person"
)

func (qs *personQueryService) GetPerson(ctx context.Context, query GetPersonQuery) (*projections.Person, error) {
	const q = `SELECT * FROM persons WHERE id=$1`

	person := projections.Person{}
	err := qs.Tx.GetDB(ctx).Get(&person, q, query.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repoperson.ErrNotFound
		}
		return nil, err
	}

	return &person, nil
}

type GetPersonQuery struct {
	Id uuid.UUID
}
