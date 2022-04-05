package memory

import (
	"testing"
	"time"

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
	ds.Clear(imagedatastore.Image{Image: "foo"})
	_, err = ds.Get(imagedatastore.Image{Image: "foo"})
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestTTL(t *testing.T) {
	now := time.Now()

	ds := New(WithNow(func() time.Time { return now }))
	ds.Set(imagedatastore.Image{Image: "foo"}, "bar")
	ret, err := ds.Get(imagedatastore.Image{Image: "foo"})
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}
	if ret != "bar" {
		t.Error("expected bar, got ", ret)
	}
	now = now.Add(2 * time.Hour)
	ret, err = ds.Get(imagedatastore.Image{Image: "foo"})
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestShortTTL(t *testing.T) {
	now := time.Now()

	ds := New(WithNow(func() time.Time { return now }), WithTTL(time.Minute))
	ds.Set(imagedatastore.Image{Image: "foo"}, "bar")
	ret, err := ds.Get(imagedatastore.Image{Image: "foo"})
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}
	if ret != "bar" {
		t.Error("expected bar, got ", ret)
	}
	now = now.Add(2 * time.Minute)
	ret, err = ds.Get(imagedatastore.Image{Image: "foo"})
	if err == nil {
		t.Error("expected error, got nil")
	}
}
