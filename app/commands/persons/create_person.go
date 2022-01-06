package persons

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"personService/domain"
)

func (cs *personCommandService) CreatePerson(ctx context.Context, command CreatePersonCommand) (uuid.UUID, error) {
	personId := uuid.New()
	person, err := domain.NewPerson(personId, command.FirstName, command.LastName)
	if err != nil {
		return uuid.Nil, nil
	}

	err = cs.PersonRepo.Insert(ctx, person)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "create person command: inserting person DB error")
	}

	return personId, nil
}

type CreatePersonCommand struct {
	FirstName string
	LastName  string
}
