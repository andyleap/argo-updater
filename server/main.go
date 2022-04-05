package server

import (
	"net/http"

	"github.com/andyleap/argo-updater/server/imagedatastore"
)

type Server struct {
	ds imagedatastore.ImageDataStore
}

func New(ds imagedatastore.ImageDataStore) *Server {
	return &Server{ds}
}

func (s *Server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	image := imagedatastore.Image{
		Image: q.Get("image"),
		Tag:   q.Get("tag"),
	}
	if image.Tag == "" {
		image.Tag = "latest"
	}
	if q.Get("cache") != "" {
		s.ds.Set(image, "")
	}
	digest, err := s.ds.Get(image)
	if err != nil {
		http.Error(rw, "not found", http.StatusNotFound)
		return
	}
	rw.Write([]byte(digest))
}
