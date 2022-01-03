package persons

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (cs *personCommandService) RenamePerson(ctx context.Context, command RenamePersonCommand) error {
	person, err := cs.personRepo.Get(ctx, command.Id)
	if err != nil {
		return errors.Wrap(err, "getting person from db error")
	}

	err = person.ChangeName(command.FirstName, command.LastName)
	if err != nil {
		return err
	}

	return cs.personRepo.Update(ctx, person)
}

type RenamePersonCommand struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
}
