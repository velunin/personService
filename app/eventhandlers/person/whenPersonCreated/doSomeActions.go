package whenPersonCreated

import (
	"context"
	"go.uber.org/fx"
	"personService/app"
)

type DoSomeActions interface {
	app.EventHandler
}

type doSomeActionsHandler struct {
	DoSomeActionsParams
}

type DoSomeActionsParams struct {
	fx.In
}

func (h *doSomeActionsHandler) Handle(ctx context.Context, event interface{}) error {
	// do some actions when person created
	return nil
}

func NewDoSomeActionsHandler(params DoSomeActionsParams) DoSomeActions {
	return &doSomeActionsHandler{params}
}
