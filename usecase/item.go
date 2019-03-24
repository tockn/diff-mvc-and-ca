package usecase

import (
	"github.com/tockn/diff-mvc-and-ca/domain/repository"
	"github.com/tockn/diff-mvc-and-ca/usecase/input"
	"github.com/tockn/diff-mvc-and-ca/usecase/output"
)

type Item interface {
	Get(ipt *input.GetItem) (*output.Item, error)
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
	oItem := &output.Item{
		ID:   ipt.ID,
		Name: item.Name,
		Rate: item.Rate,
	}
	return oItem, nil
}
