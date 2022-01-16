package commands

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"personService/modules/persons/domain"
)

func (cs *personCommandService) CreatePerson(ctx context.Context, command CreatePersonCommand) (uuid.UUID, error) {
	personId := uuid.New()
	person, err := domain.NewPerson(personId, command.FirstName, command.LastName)
	if err != nil {
		return uuid.Nil, err
	}

	txErr := cs.Tx.ExecInTran(ctx, func(ctx context.Context) error {
		err = cs.PersonRepo.Insert(ctx, person)
		if err != nil {
			return errors.Wrap(err, "create person command: inserting person DB error")
		}

		return nil
	})

	if txErr != nil {
		return uuid.Nil, txErr
	}

	return personId, nil
}

type CreatePersonCommand struct {
	FirstName string
	LastName  string
}
