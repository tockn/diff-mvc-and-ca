package usecase

import (
	"github.com/tockn/diff-mvc-and-ca/domain/repository"
	"github.com/tockn/diff-mvc-and-ca/domain/service"
	"github.com/tockn/diff-mvc-and-ca/usecase/input"
	"github.com/tockn/diff-mvc-and-ca/usecase/output"
)

type Review interface {
	Get(ipt *input.GetReview) (*output.Review, error)
	Post(ipt *input.PostReview) (*output.Review, error)
}

type review struct {
	reviewRepo  repository.Review
	itemRepo    repository.Item
	hashRepo    repository.Hash
	itemService service.Item
}

func NewReview(reviewRepo repository.Review, itemRepo repository.Item, hashRepo repository.Hash, itemService service.Item) Review {
	return &review{
		reviewRepo:  reviewRepo,
		itemRepo:    itemRepo,
		hashRepo:    hashRepo,
		itemService: itemService,
	}
}

func (r *review) Get(ipt *input.GetReview) (*output.Review, error) {
	id, err := r.hashRepo.Decode(ipt.ID)
	if err != nil {
		return nil, err
	}
	review, err := r.reviewRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	oReview := &output.Review{
		ID:   ipt.ID,
		Rate: review.Rate,
	}
	return oReview, nil
}

func (r *review) Post(ipt *input.PostReview) (*output.Review, error) {
	if err := ipt.Validate(); err != nil {
		return nil, err
	}

	itemID, err := r.hashRepo.Decode(ipt.ItemID)
	if err != nil {
		return nil, err
	}

	review, err := r.reviewRepo.Save(ipt.Rate, itemID)
	if err != nil {
		return nil, err
	}

	rate, err := r.itemService.CalcItemRateByID(itemID)
	if err != nil {
		return nil, err
	}

	_, err = r.itemRepo.UpdateRateByID(itemID, rate)
	if err != nil {
		return nil, err
	}

	id, err := r.hashRepo.Encode(review.ID)
	if err != nil {
		return nil, err
	}
	oReview := &output.Review{
		ID:   id,
		Rate: review.Rate,
	}
	return oReview, nil
}
