package repository

import "context"

type Hash interface {
	Encode(ctx context.Context, id int64) (string, error)
	Decode(ctx context.Context, idStr string) (int64, error)
}
