package whenPersonCreated

import (
	"context"
	"go.uber.org/fx"
	"personService/app"
	"personService/app/database"
	"personService/domain"
)

type CreateOutboxMessage interface {
	app.EventHandler
}

type createOutboxMessHandler struct {
	CreateOutboxMessageParams
}

type CreateOutboxMessageParams struct {
	fx.In
	Tx database.Transaction
}

func (h *createOutboxMessHandler) Handle(ctx context.Context, event interface{}) error {
	personCreated, ok := event.(domain.PersonCreatedEvent)
	if !ok {
		return nil
	}

	const query = `INSERT INTO person_created_outbox_messages (id, first_name, last_name) VALUES ($1,$2,$3)`
	_, err := h.Tx.GetDB(ctx).Exec(query, personCreated.Id, personCreated.FirstName, personCreated.LastName)
	if err != nil {
		return err
	}

	return nil
}

func NewCreateOutboxMessageHandler(params CreateOutboxMessageParams) CreateOutboxMessage {
	return &createOutboxMessHandler{params}
}
