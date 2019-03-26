package usecase

import (
	"context"

	"github.com/tockn/diff-mvc-and-ca/usecase/input"
)

type Interactor struct {
	presenter Presenter
	item      Item
	review    Review
}

func NewInteractor(pre Presenter, item Item, review Review) Interactor {
	return Interactor{
		presenter: pre,
		item:      item,
		review:    review,
	}
}

func (i *Interactor) GetItem(ctx context.Context, ipt *input.GetItem) {
	item, err := i.item.Get(ipt)
	if err != nil {
		i.presenter.ViewError(ctx, err)
		return
	}

	i.presenter.ViewItem(ctx, item)
}

func (i *Interactor) PostItem(ctx context.Context, ipt *input.PostItem) {
	item, err := i.item.Post(ipt)
	if err != nil {
		i.presenter.ViewError(ctx, err)
		return
	}

	i.presenter.ViewPostItem(ctx, item)
}

func (i *Interactor) GetReview(ctx context.Context, ipt *input.GetReview) {
	review, err := i.review.Get(ipt)
	if err != nil {
		i.presenter.ViewError(ctx, err)
		return
	}

	i.presenter.ViewReview(ctx, review)
}

func (i *Interactor) PostReview(ctx context.Context, ipt *input.PostReview) {
	review, err := i.review.Post(ipt)
	if err != nil {
		i.presenter.ViewError(ctx, err)
		return
	}

	i.presenter.ViewPostReview(ctx, review)
}
