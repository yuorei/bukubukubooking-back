package infrastructure

import (
	"context"

	"github.com/yuorei/bukubukubooking-back/src/domain"
)

func (i *Infrastructure) GetUserFromDB(ctx context.Context, id string) (*domain.DBUser, error) {
	return nil, nil
}

func (i *Infrastructure) InsertUser(ctx context.Context, user *domain.DBUser) (*domain.DBUser, error) {
	return user, nil
}
