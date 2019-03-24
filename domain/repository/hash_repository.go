package repository

type Hash interface {
	Encode(id int64) (string, error)
	Decode(idStr string) (int64, error)
}
