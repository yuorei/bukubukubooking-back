package application

import (
	"github.com/yuorei/bukubukubooking-back/src/application/port"
)

type UserUseCase struct {
	userRepository port.UserRepository
}

func NewUserUseCase(userRepository port.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
	}
}
