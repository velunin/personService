package persons

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (cs *personCommandService) BlockPerson(ctx context.Context, command BlockPersonCommand) error {
	return cs.Tx.ExecInTran(ctx, func(ctx context.Context) error {
		person, err := cs.PersonRepo.Get(ctx, command.Id)
		if err != nil {
			return errors.Wrap(err, "block person command: getting person from DB error")
		}

		err = person.Block()
		if err != nil {
			return err
		}

		err = cs.PersonRepo.Update(ctx, person)
		if err != nil {
			return errors.Wrap(err, "block person command: updating person DB error")
		}

		return nil
	})
}

type BlockPersonCommand struct {
	Id uuid.UUID
}
