package grokking_algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name    string
		i       int
		slice   []int
		want    int
		wantErr error
	}{
		{"not found", 10, []int{1, 2, 3, 4, 5, 5}, 0, NotFound},
		{"empty slice", 0, []int{}, 0, NotFound},
		{"first element", 1, []int{1, 2, 3, 4}, 0, nil},
		{"second element", 2, []int{1, 2, 3, 4}, 1, nil},
		{"last element", 4, []int{1, 2, 3, 4}, 3, nil},
		{"only element", 4, []int{4}, 0, nil},
		{"two elements", 3, []int{1, 2, 3, 3, 4}, 2, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinarySearch(tt.i, tt.slice)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
