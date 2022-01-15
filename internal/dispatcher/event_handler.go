package dispatcher

import (
	"context"
	"go.uber.org/fx"
)

type EventHandler interface {
	Handle(ctx context.Context, event interface{}) error
	GetEventType() interface{}
}

type EventHandlerGroup struct {
	fx.Out

	Handler EventHandler `group:"event_handlers"`
}
