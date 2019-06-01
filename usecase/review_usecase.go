package usecase

import (
	"context"

	"github.com/tockn/diff-mvc-and-ca/domain/repository"
	"github.com/tockn/diff-mvc-and-ca/domain/service"
	"github.com/tockn/diff-mvc-and-ca/usecase/input"
	"github.com/tockn/diff-mvc-and-ca/usecase/output"
)

type Review interface {
	Get(ctx context.Context, ipt *input.GetReview) (*output.Review, error)
	Post(ctx context.Context, ipt *input.PostReview) (*output.Review, error)
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

func (r *review) Get(ctx context.Context, ipt *input.GetReview) (*output.Review, error) {
	id, err := r.hashRepo.Decode(ctx, ipt.ID)
	if err != nil {
		return nil, err
	}
	review, err := r.reviewRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	oReview := &output.Review{
		ID:   ipt.ID,
		Rate: review.Rate,
	}
	return oReview, nil
}

func (r *review) Post(ctx context.Context, ipt *input.PostReview) (*output.Review, error) {

	// バリデーション
	if err := ipt.Validate(); err != nil {
		return nil, err
	}

	// 入力されたレビューデータの商品ハッシュIDを数値IDへ変換
	itemID, err := r.hashRepo.Decode(ctx, ipt.ItemID)
	if err != nil {
		return nil, err
	}

	// 入力されたレビューデータを永続化
	review, err := r.reviewRepo.Save(ctx, ipt.Rate, itemID)
	if err != nil {
		return nil, err
	}

	// 永続化したレビューデータの数値IDをハッシュIDへ変換
	id, err := r.hashRepo.Encode(ctx, review.ID)
	if err != nil {
		return nil, err
	}

	// レビューされた商品のレート更新
	rate, err := r.itemService.CalcItemRateByID(ctx, itemID)
	if err != nil {
		return nil, err
	}

	_, err = r.itemRepo.UpdateRateByID(ctx, itemID, rate)
	if err != nil {
		return nil, err
	}

	// レビュー情報を出力
	oReview := &output.Review{
		ID:   id,
		Rate: review.Rate,
	}
	return oReview, nil
}
