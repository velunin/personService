package persons

import (
	"context"
)

func (cs *personCommandService) CreatePerson(ctx context.Context, command CreatePersonCommand) error {
	panic("")
}

type CreatePersonCommand struct {
	FirstName string
	LastName  string
}
