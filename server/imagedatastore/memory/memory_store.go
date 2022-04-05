package memory

import (
	"fmt"

	"github.com/andyleap/argo-updater/server/imagedatastore"
)

type memoryImageDataStore struct {
	data map[imagedatastore.Image]imagedatastore.Digest
}

func New() imagedatastore.ImageDataStore {
	return &memoryImageDataStore{
		data: map[imagedatastore.Image]imagedatastore.Digest{},
	}
}

func (m *memoryImageDataStore) Get(image imagedatastore.Image) (imagedatastore.Digest, error) {
	d, ok := m.data[image]
	if !ok {
		return "", fmt.Errorf("image not found")
	}
	return d, nil
}

func (m *memoryImageDataStore) Set(image imagedatastore.Image, digest imagedatastore.Digest) error {
	if digest == "" {
		delete(m.data, image)
		return nil
	}
	m.data[image] = digest
	return nil
}
