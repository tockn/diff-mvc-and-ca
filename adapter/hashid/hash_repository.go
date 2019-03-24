package hashid

import (
	"errors"

	"github.com/speps/go-hashids"
	"github.com/tockn/diff-mvc-and-ca/domain/repository"
)

// Implement repository.Hash
type hash struct {
	*hashids.HashID
}

func NewHash(h *hashids.HashID) repository.Hash {
	return &hash{
		h,
	}
}

func (h *hash) Encode(id int64) (string, error) {
	return h.EncodeInt64([]int64{id})
}

func (h *hash) Decode(idStr string) (int64, error) {
	nums, err := h.DecodeInt64WithError(idStr)
	if len(nums) != 1 {
		return 0, errors.New("decode error")
	}
	return nums[0], err
}
