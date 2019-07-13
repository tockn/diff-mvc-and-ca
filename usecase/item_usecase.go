package usecase

import (
	"context"

	"github.com/tockn/diff-mvc-and-ca/domain/repository"
	"github.com/tockn/diff-mvc-and-ca/usecase/input"
	"github.com/tockn/diff-mvc-and-ca/usecase/output"
)

type Item interface {
	Get(ctx context.Context, ipt *input.GetItem) (*output.Item, error)
	Post(ctx context.Context, ipt *input.PostItem) (*output.Item, error)
}

type item struct {
	itemRepo repository.Item
	hashRepo repository.Hash
}

func NewItem(itemRepo repository.Item, hashRepo repository.Hash) Item {
	return &item{
		itemRepo: itemRepo,
		hashRepo: hashRepo,
	}
}

func (i *item) Get(ctx context.Context, ipt *input.GetItem) (*output.Item, error) {
	id, err := i.hashRepo.Decode(ctx, ipt.ID)
	if err != nil {
		return nil, err
	}
	item, err := i.itemRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	ranking, err := i.itemRepo.GetRankingByID(ctx, id)
	if err != nil {
		return nil, err
	}
	oItem := &output.Item{
		ID:      ipt.ID,
		Name:    item.Name,
		Rate:    item.Rate,
		Ranking: ranking,
	}
	return oItem, nil
}

func (i *item) Post(ctx context.Context, ipt *input.PostItem) (*output.Item, error) {
	if err := ipt.Validate(); err != nil {
		return nil, err
	}
	item, err := i.itemRepo.Save(ctx, ipt.Name)
	if err != nil {
		return nil, err
	}
	id, err := i.hashRepo.Encode(ctx, item.ID)
	if err != nil {
		return nil, err
	}
	ranking, err := i.itemRepo.GetRankingByID(ctx, item.ID)
	if err != nil {
		return nil, err
	}
	oItem := &output.Item{
		ID:      id,
		Name:    item.Name,
		Rate:    item.Rate,
		Ranking: ranking,
	}
	return oItem, nil
}
