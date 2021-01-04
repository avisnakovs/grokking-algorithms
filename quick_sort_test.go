package grokking_algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{"nil", nil, nil},
		{"empty", []int{}, []int{}},
		{"one element", []int{1}, []int{1}},
		{"two element", []int{2, 1}, []int{1, 2}},
		{"three element", []int{2, 3, 1}, []int{1, 2, 3}},
		{"repeating elements", []int{2, 3, 4, 7, 5, 5, 4, 1}, []int{1, 2, 3, 4, 4, 5, 5, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := QuickSort(tt.slice)
			assert.Equal(t, tt.want, got)
		})
	}
}
