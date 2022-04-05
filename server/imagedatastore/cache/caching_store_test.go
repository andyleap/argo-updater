package cache

import (
	"testing"

	"github.com/andyleap/argo-updater/server/imagedatastore"
	"github.com/andyleap/argo-updater/server/imagedatastore/memory"
)

func TestCachingStore(t *testing.T) {
	p := &mockStore{ImageDataStore: memory.New()}
	c := memory.New()
	cs := New(p, c)

	p.Set(imagedatastore.Image{Image: "foo"}, "bar")
	ret, err := cs.Get(imagedatastore.Image{Image: "foo"})
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}
	if ret != "bar" {
		t.Error("expected bar, got ", ret)
	}
	ret, err = c.Get(imagedatastore.Image{Image: "foo"})
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}
	if ret != "bar" {
		t.Error("expected bar, got ", ret)
	}
	ret, err = cs.Get(imagedatastore.Image{Image: "foo"})
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}
	if ret != "bar" {
		t.Error("expected bar, got ", ret)
	}
	if len(p.gets) != 1 {
		t.Error("expected 1 get, got ", len(p.gets))
	}
}

type mockStore struct {
	imagedatastore.ImageDataStore
	gets []imagedatastore.Image
}

func (ms *mockStore) Get(image imagedatastore.Image) (imagedatastore.Digest, error) {
	ms.gets = append(ms.gets, image)
	return ms.ImageDataStore.Get(image)
}
