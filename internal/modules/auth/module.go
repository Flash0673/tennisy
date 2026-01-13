package auth

import "tennisy.com/mvp/internal/modules/auth/action"

type Module struct {
	Actions *action.Aggregator
}

func New() *Module {
	return &Module{
		Actions: action.New(),
	}
}
