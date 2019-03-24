package output

type Item struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Rate    float64 `json:"rate"`
	Ranking int64   `json:"ranking"`
}
