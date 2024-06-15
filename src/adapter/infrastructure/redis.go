package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/yuorei/bukubukubooking-back/src/domain"
)

func getFromRedis(ctx context.Context, client *redis.Client, key string, data any) (bool, error) {
	bytes, err := client.Get(ctx, key).Bytes()
	if err != nil {
		if err.Error() == "redis: nil" {
			return false, nil
		}

		return false, err
	}

	switch v := data.(type) {
	case *[]*domain.Book:
		err = json.Unmarshal(bytes, v)
		if err != nil {
			return false, err
		}
	default:
		return false, fmt.Errorf("invalid type")
	}

	return true, nil
}

func setToRedis(ctx context.Context, client *redis.Client, key string, expiration time.Duration, value any) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	var books []*domain.Book
	err = json.Unmarshal(bytes, &books)
	if err != nil {
		return err
	}

	return client.Set(ctx, key, bytes, expiration).Err()
}
