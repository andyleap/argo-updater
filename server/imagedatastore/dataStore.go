package imagedatastore

type Image struct {
	Image string
	Tag   string
}

type Digest string

type ImageDataStore interface {
	Get(Image) (Digest, error)
	Set(Image, Digest) error
}
