package memory

import (
	"fmt"
	"time"

	"github.com/andyleap/argo-updater/server/imagedatastore"
)

type entry struct {
	digest imagedatastore.Digest
	expiry time.Time
}

type memoryImageDataStore struct {
	data map[imagedatastore.Image]entry
	now  func() time.Time
	ttl  time.Duration
}

type Option func(*memoryImageDataStore)

func WithTTL(ttl time.Duration) Option {
	return func(m *memoryImageDataStore) {
		m.ttl = ttl
	}
}

func WithNow(now func() time.Time) Option {
	return func(m *memoryImageDataStore) {
		m.now = now
	}
}

func New(options ...Option) imagedatastore.ImageDataStore {
	ds := &memoryImageDataStore{
		data: map[imagedatastore.Image]entry{},
		now:  time.Now,
		ttl:  time.Hour,
	}
	for _, option := range options {
		option(ds)
	}
	return ds
}

func (m *memoryImageDataStore) Get(image imagedatastore.Image) (imagedatastore.Digest, error) {
	d, ok := m.data[image]
	if !ok {
		return "", fmt.Errorf("image not found")
	}
	if d.expiry.Before(m.now()) {
		delete(m.data, image)
		return "", fmt.Errorf("image not found")
	}
	return d.digest, nil
}

func (m *memoryImageDataStore) Set(image imagedatastore.Image, digest imagedatastore.Digest) error {
	m.data[image] = entry{
		digest: digest,
		expiry: m.now().Add(m.ttl),
	}
	return nil
}

func (m *memoryImageDataStore) Clear(image imagedatastore.Image) error {
	delete(m.data, image)
	return nil
}
