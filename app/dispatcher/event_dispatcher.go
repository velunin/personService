package dispatcher

import (
	"context"
	"errors"
	"go.uber.org/fx"
	"personService/app"
	"personService/app/eventhandlers/person/whenPersonCreated"
	"personService/domain"
	"reflect"
)

type Dispatcher interface {
	Dispatch(ctx context.Context, event interface{}) error
}
type dispatcher struct {
	Params
}

type Params struct {
	fx.In

	WhenPersonCreatedCreateOutboxMessageHandler whenPersonCreated.CreateOutboxMessage
	WhenPersonCreatedDoSomeActionsHandler       whenPersonCreated.DoSomeActions
}

var handlersMap = map[string][]app.EventHandler{}

func (d *dispatcher) Dispatch(ctx context.Context, event interface{}) error {
	if event == nil {
		return errors.New("event is nil")
	}

	typeName := reflect.TypeOf(event).String()
	if h, ok := handlersMap[typeName]; ok {
		for _, handler := range h {
			err := handler.Handle(ctx, event)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func New(params Params) Dispatcher {
	d := &dispatcher{params}

	setupHandlers(d)

	return d
}

func setupHandlers(d *dispatcher) {
	registerHandler(domain.PersonCreated{},
		d.WhenPersonCreatedCreateOutboxMessageHandler,
		d.WhenPersonCreatedDoSomeActionsHandler)
}

func registerHandler(event interface{}, handlers ...app.EventHandler) {
	typeName := reflect.TypeOf(event).String()
	eventHandlers := handlersMap[typeName]
	eventHandlers = append(eventHandlers, handlers...)
	handlersMap[typeName] = eventHandlers
}
