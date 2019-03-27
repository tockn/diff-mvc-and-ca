package response

import "github.com/tockn/diff-mvc-and-ca/usecase/output"

type Review struct {
	ID   string  `json:"id"`
	Rate float64 `json:"rate"`
}

func NewReviewFromOutput(or *output.Review) *Review {
	return &Review{
		ID:   or.ID,
		Rate: or.Rate,
	}
}
