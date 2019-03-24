package registry

import (
	"github.com/tockn/diff-mvc-and-ca/adapter/datastore"
	"github.com/tockn/diff-mvc-and-ca/domain/repository"
	"github.com/tockn/diff-mvc-and-ca/domain/service"
	"github.com/tockn/diff-mvc-and-ca/usecase"
)

func (r *registry) NewItemUsecase() usecase.Item {
	return usecase.NewItem(r.NewItemRepository(), r.NewHashRepository())
}

func (r *registry) NewItemService() service.Item {
	return service.NewItem(r.NewReviewRepository())
}

func (r *registry) NewItemRepository() repository.Item {
	return datastore.NewItem(r.db)
}
