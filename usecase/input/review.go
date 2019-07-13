package input

import "errors"

type GetReview struct {
	ID     string
	ItemID string
}

type PostReview struct {
	Rate   float64
	ItemID string
}

func (r *PostReview) Validate() error {
	if r.Rate < 1 || 5 < r.Rate {
		return errors.New("[PostReview] Rate validate error")
	}
	return nil
}
