package action

import "tennisy.com/mvp/internal/modules/auth/action/register"

type Aggregator struct {
	// Сюда собираем экшены
	Register *register.Action
}

func New() *Aggregator {
	return &Aggregator{
		Register: register.New(),
	}
}
