package docker

import (
	"github.com/andyleap/argo-updater/server/imagedatastore"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

type dockerStore struct {
	options []remote.Option
}

func New(options ...remote.Option) imagedatastore.ImageDataStore {
	return &dockerStore{
		options: options,
	}
}

func (ds *dockerStore) Get(image imagedatastore.Image) (imagedatastore.Digest, error) {
	ref, err := name.ParseReference(image.Image + ":" + image.Tag)
	if err != nil {
		return "", err
	}
	d, err := remote.Head(ref, ds.options...)
	if err != nil {
		return "", nil
	}
	return imagedatastore.Digest(d.Digest.String()), nil
}

func (ds *dockerStore) Set(_ imagedatastore.Image, _ imagedatastore.Digest) error {
	return nil
}

func (ds *dockerStore) Clear(_ imagedatastore.Image) error {
	return nil
}
