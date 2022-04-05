package config

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestConfigParse(t *testing.T) {
	rawConf := `
imagedatastore:
  cache:
    primary:
      docker: {}
    cache:
      memory: {}
`
	conf := Config{}
	err := yaml.Unmarshal([]byte(rawConf), &conf)
	if err != nil {
		t.Fatal(err)
	}
	if conf.ImageDataStore.Cache == nil {
		t.Fatal("Cache is nil")
	}
	if conf.ImageDataStore.Cache.Primary.Docker == nil {
		t.Fatal("Primary.Docker is nil")
	}
	if conf.ImageDataStore.Cache.Cache.Memory == nil {
		t.Fatal("Cache.Memory is nil")
	}
}
