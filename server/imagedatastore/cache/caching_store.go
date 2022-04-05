package cache

import "github.com/andyleap/argo-updater/server/imagedatastore"

type cachingStore struct {
	primary imagedatastore.ImageDataStore
	cache   imagedatastore.ImageDataStore
}

func New(primary, cache imagedatastore.ImageDataStore) imagedatastore.ImageDataStore {
	return &cachingStore{primary, cache}
}

func (cs *cachingStore) Get(image imagedatastore.Image) (imagedatastore.Digest, error) {
	d, err := cs.cache.Get(image)
	if err == nil {
		return d, nil
	}
	d, err = cs.primary.Get(image)
	if err != nil {
		return "", err
	}
	cs.cache.Set(image, d)
	return d, nil
}

func (cs *cachingStore) Set(image imagedatastore.Image, digest imagedatastore.Digest) error {
	cs.cache.Set(image, digest)
	return cs.primary.Set(image, digest)
}

func (cs *cachingStore) Clear(image imagedatastore.Image) error {
	cs.cache.Clear(image)
	return cs.primary.Clear(image)
}
