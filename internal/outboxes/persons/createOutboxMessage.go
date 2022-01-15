package persons

import (
	"context"
	"go.uber.org/fx"
	"personService/internal/database"
	"personService/internal/dispatcher"
	"personService/modules/persons/domain"
)

type CreateOutboxMessage interface {
	dispatcher.EventHandler
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

func (h *createOutboxMessHandler) GetEventType() interface{} {
	return domain.PersonCreatedEvent{}
}

func NewCreateOutboxMessageHandler(params CreateOutboxMessageParams) dispatcher.EventHandlerGroup {
	return dispatcher.EventHandlerGroup{Handler: &createOutboxMessHandler{params}}
}
