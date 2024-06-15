package infrastructure

import (
	"github.com/redis/go-redis/v9"
	"github.com/yuorei/bukubukubooking-back/src/driver/db"
	r "github.com/yuorei/bukubukubooking-back/src/driver/redis"
)

type Infrastructure struct {
	redis *redis.Client
	dbc   *db.DB
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{
		redis: r.ConnectRedis(),
		dbc:   db.ConnectDB(),
	}
}
