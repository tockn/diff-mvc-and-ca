package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReview_Insert(t *testing.T) {

}

func TestReview_Validate(t *testing.T) {
	tests := []struct {
		review Review
		expect bool
	}{
		{Review{Rate: 0}, false},
		{Review{Rate: 1}, true},
		{Review{Rate: 5}, true},
		{Review{Rate: 6}, false},
	}
	for _, tt := range tests {
		got := tt.review.Validate()
		if !cmp.Equal(tt.expect, got) {
			t.Fatalf("expect: %v, got: %v", tt.expect, got)
		}
	}
}

func TestCalculateRate(t *testing.T) {
	tests := []struct {
		sum    int64
		count  int64
		expect float64
	}{
		{0, 0, 0},
		{1, 1, 1},
		{5, 2, 2.5},
	}
	for _, tt := range tests {
		got := CalculateRate(tt.sum, tt.count)
		if !cmp.Equal(tt.expect, got) {
			t.Fatalf("expect: %f, got: %f", tt.expect, got)
		}
	}
}
