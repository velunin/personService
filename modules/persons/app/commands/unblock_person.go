package commands

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (cs *personCommandService) UnblockPerson(ctx context.Context, command UnblockPersonCommand) error {
	return cs.Tx.ExecInTran(ctx, func(ctx context.Context) error {
		person, err := cs.PersonRepo.Get(ctx, command.Id)
		if err != nil {
			return errors.Wrap(err, "unblock person command: getting person from DB error")
		}

		err = person.Unblock()
		if err != nil {
			return err
		}

		err = cs.PersonRepo.Update(ctx, person)
		if err != nil {
			return errors.Wrap(err, "unblock person command: updating person DB error")
		}

		return nil
	})
}

type UnblockPersonCommand struct {
	Id uuid.UUID
}
