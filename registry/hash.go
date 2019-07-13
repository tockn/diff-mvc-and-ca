package registry

import (
	"github.com/tockn/diff-mvc-and-ca/adapter/hashid"
	"github.com/tockn/diff-mvc-and-ca/domain/repository"
)

func (r *registry) NewHashRepository() repository.Hash {
	return hashid.NewHash(r.hash)
}
