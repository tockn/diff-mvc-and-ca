package usecase

import (
	"errors"

	"github.com/tockn/diff-mvc-and-ca/domain/repository"
	"github.com/tockn/diff-mvc-and-ca/usecase/input"
	"github.com/tockn/diff-mvc-and-ca/usecase/output"
)

type Item interface {
	Get(ipt *input.GetItem) (*output.Item, error)
	Post(ipt *input.PostItem) (*output.Item, error)
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

func (i *item) Get(ipt *input.GetItem) (*output.Item, error) {
	id, err := i.hashRepo.Decode(ipt.ID)
	if err != nil {
		return nil, err
	}
	item, err := i.itemRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	ranking, err := i.itemRepo.GetRankingByID(id)
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

func (i *item) Post(ipt *input.PostItem) (*output.Item, error) {
	if !ipt.Validate() {
		return nil, errors.New("validation error")
	}
	item, err := i.itemRepo.Save(ipt.Name)
	if err != nil {
		return nil, err
	}
	id, err := i.hashRepo.Encode(item.ID)
	if err != nil {
		return nil, err
	}
	ranking, err := i.itemRepo.GetRankingByID(item.ID)
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
