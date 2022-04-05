package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/andyleap/argo-updater/server/imagedatastore"
	"github.com/go-redis/redis/v8"
)

type redisStore struct {
	client *redis.Client
	ttl    time.Duration
}

func New(client *redis.Client, ttl time.Duration) imagedatastore.ImageDataStore {
	return &redisStore{
		client: client,
		ttl:    ttl,
	}
}

func (rs *redisStore) Get(image imagedatastore.Image) (imagedatastore.Digest, error) {
	d, err := rs.client.Get(context.Background(), fmt.Sprintf("%s:%s", image.Image, image.Tag)).Result()
	return imagedatastore.Digest(d), err
}

func (rs *redisStore) Set(image imagedatastore.Image, digest imagedatastore.Digest) error {
	return rs.client.Set(context.Background(), fmt.Sprintf("%s:%s", image.Image, image.Tag), digest, rs.ttl).Err()
}

func (rs *redisStore) Clear(image imagedatastore.Image) error {
	return rs.client.Del(context.Background(), fmt.Sprintf("%s:%s", image.Image, image.Tag)).Err()
}
