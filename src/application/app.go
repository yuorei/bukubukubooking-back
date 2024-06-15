package application

import (
	"github.com/yuorei/bukubukubooking-back/src/adapter/infrastructure"
)

type Application struct {
	User *UserUseCase
}

func NewApplication(infra *infrastructure.Infrastructure) *Application {
	userUseCase := NewUserUseCase(infra)

	return &Application{
		User: userUseCase,
	}
}
