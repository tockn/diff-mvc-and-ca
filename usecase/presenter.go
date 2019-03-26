package usecase

import (
	"context"

	"github.com/tockn/diff-mvc-and-ca/usecase/output"
)

type Presenter interface {
	ViewItem(ctx context.Context, item *output.Item)
	ViewPostItem(ctx context.Context, item *output.Item)
	ViewReview(ctx context.Context, review *output.Review)
	ViewPostReview(ctx context.Context, review *output.Review)
	ViewError(ctx context.Context, err error)
}
