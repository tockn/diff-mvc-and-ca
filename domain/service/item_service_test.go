package service

import (
	"testing"

	mock_repository "github.com/tockn/diff-mvc-and-ca/adapter/mock/repository"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestItem_CalcItemRateByID(t *testing.T) {
	tests := []struct {
		id       int64
		count    int64
		sum      int64
		expected float64
	}{
		{1, 10, 30, 3},
		{1, 0, 0, 0},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReviewRepository := mock_repository.NewMockReview(ctrl)

	itemService := NewItem(mockReviewRepository)

	for _, tt := range tests {
		mockReviewRepository.EXPECT().CountByItemID(gomock.Eq(tt.id)).Return(tt.count, nil)
		mockReviewRepository.EXPECT().SumOfRateByItemID(gomock.Eq(tt.id)).Return(tt.sum, nil)
		actual, err := itemService.CalcItemRateByID(tt.id)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, tt.expected, actual)
	}
}
