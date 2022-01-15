package queries

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	repoperson "personService/modules/persons/app"
	"personService/modules/persons/domain"
)

func (qs *personQueryService) GetPerson(ctx context.Context, query GetPersonQuery) (*Person, error) {
	const q = `SELECT * FROM persons WHERE id=$1`

	person := Person{}
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

type Person struct {
	domain.PersonState
}
