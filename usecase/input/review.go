package input

type GetReview struct {
	ID     string `json:"id"`
	ItemID string `json:"item_id"`
}

type PostReview struct {
	Rate   float64 `json:"rate"`
	ItemID string  `json:"item_id"`
}

func (r *PostReview) ValidatePostReview() bool {
	if r.Rate < 1 || 5 < r.Rate {
		return false
	}
	return true
}
