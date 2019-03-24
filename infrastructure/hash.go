package infrastructure

import "github.com/speps/go-hashids"

func NewHash() (*hashids.HashID, error) {
	h := hashids.NewData()
	h.MinLength = 5
	return hashids.NewWithData(h)
}
