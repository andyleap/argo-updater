package main

import (
	"log"
	"net/http"
	"os"

	"github.com/andyleap/argo-updater/server"
	"github.com/andyleap/argo-updater/server/config"
	"github.com/jessevdk/go-flags"
	"gopkg.in/yaml.v3"
)

func main() {
	var options struct {
		ConfigFile string `short:"c" long:"config" description:"Path to config file"`
	}
	_, err := flags.Parse(&options)
	if err != nil {
		log.Fatal(err)
	}
	conf := config.Config{}
	if options.ConfigFile != "" {
		f, err := os.Open(options.ConfigFile)
		if err != nil {
			log.Fatal(err)
		}
		err = yaml.NewDecoder(f).Decode(&conf)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		conf = config.DefaultConfig
	}
	ds := conf.ImageDataStore.Get()
	s := server.New(ds)
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
