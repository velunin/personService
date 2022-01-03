package persons

import (
	"context"
	"go.uber.org/fx"
	"personService/app/repositories/person"
)

type PersonCommandService interface {
	CreatePerson(ctx context.Context, command CreatePersonCommand) error
	RenamePerson(ctx context.Context, command RenamePersonCommand) error
}

type personCommandService struct {
	fx.In
	personRepo person.Repository
}

func NewPersonCommandService() PersonCommandService {
	return &personCommandService{}
}
