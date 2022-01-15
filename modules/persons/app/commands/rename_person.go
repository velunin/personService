package commands

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (cs *personCommandService) RenamePerson(ctx context.Context, command RenamePersonCommand) error {
	return cs.Tx.ExecInTran(ctx, func(ctx context.Context) error {
		person, err := cs.PersonRepo.Get(ctx, command.Id)
		if err != nil {
			return errors.Wrap(err, "rename person command: getting person from db error")
		}

		err = person.ChangeName(command.FirstName, command.LastName)
		if err != nil {
			return err
		}

		err = cs.PersonRepo.Update(ctx, person)
		if err != nil {
			return errors.Wrap(err, "rename person command: updating person DB error")
		}

		return nil
	})
}

type RenamePersonCommand struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
}
