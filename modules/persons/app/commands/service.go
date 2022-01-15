package commands

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/fx"
	"personService/internal/database"
	"personService/modules/persons/app"
)

type PersonCommandService interface {
	CreatePerson(ctx context.Context, command CreatePersonCommand) (uuid.UUID, error)
	RenamePerson(ctx context.Context, command RenamePersonCommand) error
	BlockPerson(ctx context.Context, command BlockPersonCommand) error
	UnblockPerson(ctx context.Context, command UnblockPersonCommand) error
}

type CsParams struct {
	fx.In
	PersonRepo app.Repository
	Tx         database.Transaction
}

type personCommandService struct {
	CsParams
}

func NewPersonCommandService(params CsParams) PersonCommandService {
	return &personCommandService{params}
}
