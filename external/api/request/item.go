package request

import "github.com/tockn/diff-mvc-and-ca/usecase/input"

type GetItem struct {
	ID string `json:"id"`
}

func (gi *GetItem) ToInput() *input.GetItem {
	return &input.GetItem{
		ID: gi.ID,
	}
}

type PostItem struct {
	Name string `json:"name"`
}

func (pi *PostItem) ToInput() *input.PostItem {
	return &input.PostItem{
		Name: pi.Name,
	}
}
