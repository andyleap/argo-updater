package config

import (
	"time"

	"github.com/andyleap/argo-updater/server/imagedatastore"
	"github.com/andyleap/argo-updater/server/imagedatastore/cache"
	"github.com/andyleap/argo-updater/server/imagedatastore/docker"
	"github.com/andyleap/argo-updater/server/imagedatastore/memory"
	redisds "github.com/andyleap/argo-updater/server/imagedatastore/redis"
	"github.com/go-redis/redis/v8"
)

type Config struct {
	ImageDataStore ImageDataStore
}

var DefaultConfig = Config{
	ImageDataStore: ImageDataStore{
		Cache: &CacheDataStore{
			Primary: ImageDataStore{
				Docker: &DockerDataStore{},
			},
			Cache: ImageDataStore{
				Memory: &MemoryDataStore{},
			},
		},
	},
}

type CacheDataStore struct {
	Primary ImageDataStore
	Cache   ImageDataStore
}

type MemoryDataStore struct {
}

type RedisDataStore struct {
	redis.Options
	TTL time.Duration
}

type DockerDataStore struct {
}

type ImageDataStore struct {
	Cache  *CacheDataStore
	Memory *MemoryDataStore
	Redis  *RedisDataStore
	Docker *DockerDataStore
}

func (ids ImageDataStore) Get() imagedatastore.ImageDataStore {
	if ids.Cache != nil {
		return cache.New(ids.Cache.Primary.Get(), ids.Cache.Cache.Get())
	}
	if ids.Memory != nil {
		return memory.New()
	}
	if ids.Redis != nil {
		return redisds.New(redis.NewClient(&ids.Redis.Options), ids.Redis.TTL)
	}
	if ids.Docker != nil {
		return docker.New()
	}
	panic("Unsupported ImageDataStore")
}
