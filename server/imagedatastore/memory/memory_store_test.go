package memory

import (
	"testing"

	"github.com/andyleap/argo-updater/server/imagedatastore"
)

func TestMemoryStore(t *testing.T) {
	var ds imagedatastore.ImageDataStore
	ds = New()
	ds.Set(imagedatastore.Image{Image: "foo"}, "bar")
	ret, err := ds.Get(imagedatastore.Image{Image: "foo"})
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}
	if ret != "bar" {
		t.Error("expected bar, got ", ret)
	}
}

func TestMemoryStoreClear(t *testing.T) {
	var ds imagedatastore.ImageDataStore
	ds = New()
	ds.Set(imagedatastore.Image{Image: "foo"}, "bar")
	ret, err := ds.Get(imagedatastore.Image{Image: "foo"})
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}
	if ret != "bar" {
		t.Error("expected bar, got ", ret)
	}
	ds.Set(imagedatastore.Image{Image: "foo"}, "")
	_, err = ds.Get(imagedatastore.Image{Image: "foo"})
	if err == nil {
		t.Error("expected error, got nil")
	}
}
