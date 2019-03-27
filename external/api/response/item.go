package response

import "github.com/tockn/diff-mvc-and-ca/usecase/output"

type Item struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Rate    float64 `json:"rate"`
	Ranking int64   `json:"ranking"`
}

func NewItemFromOutput(oi *output.Item) *Item {
	return &Item{
		ID:      oi.ID,
		Name:    oi.Name,
		Rate:    oi.Rate,
		Ranking: oi.Ranking,
	}
}
