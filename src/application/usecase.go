package application

import (
	"github.com/yuorei/bukubukubooking-back/src/application/port"
)

type UseCase struct {
	port.UserInputPort
}

func NewUseCase(application *Application) *UseCase {
	return &UseCase{
		UserInputPort: application,
	}
}
