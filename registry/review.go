package registry

import (
	"github.com/tockn/diff-mvc-and-ca/adapter/datastore"
	"github.com/tockn/diff-mvc-and-ca/domain/repository"
	"github.com/tockn/diff-mvc-and-ca/usecase"
)

func (r *registry) NewReviewUsecase() usecase.Review {
	return usecase.NewReview(
		r.NewReviewRepository(),
		r.NewItemRepository(),
		r.NewHashRepository(),
		r.NewItemService(),
	)
}

func (r *registry) NewReviewRepository() repository.Review {
	return datastore.NewReview(r.db)
}
