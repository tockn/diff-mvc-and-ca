package model

import (
	"errors"

	"github.com/speps/go-hashids"
)

func EncodeID(hash *hashids.HashID, id int64) (string, error) {
	return hash.EncodeInt64([]int64{id})
}

func DecodeID(hash *hashids.HashID, idStr string) (int64, error) {
	nums, err := hash.DecodeInt64WithError(idStr)
	if len(nums) != 1 {
		return 0, errors.New("decode error")
	}
	return nums[0], err
}
