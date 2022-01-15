package dispatcher

import (
	"context"
	"errors"
	"go.uber.org/fx"
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

	Handlers []EventHandler `group:"event_handlers"`
}

var handlersMap = map[string][]EventHandler{}

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
	for _, handler := range params.Handlers {
		typeName := reflect.TypeOf(handler.GetEventType()).String()
		eventHandlers := handlersMap[typeName]
		handlersMap[typeName] = append(eventHandlers, handler)
	}

	return &dispatcher{}
}
