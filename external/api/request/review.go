package request

import "github.com/tockn/diff-mvc-and-ca/usecase/input"

type GetReview struct {
	ID     string `json:"id"`
	ItemID string `json:"item_id"`
}

func (gr *GetReview) ToInput() *input.GetReview {
	return &input.GetReview{
		ID:     gr.ID,
		ItemID: gr.ItemID,
	}
}

type PostReview struct {
	Rate   float64 `json:"rate"`
	ItemID string  `json:"item_id"`
}

func (pr *PostReview) ToInput() *input.PostReview {
	return &input.PostReview{
		Rate:   pr.Rate,
		ItemID: pr.ItemID,
	}
}
