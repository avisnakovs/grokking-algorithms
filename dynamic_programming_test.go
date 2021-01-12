package grokking_algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDynamicKnapsack(t *testing.T) {
	type args struct {
		elements  []Element
		totalSize int
	}
	tests := []struct {
		name         string
		args         args
		want         Table
		wantMaxPrice int
	}{
		{
			"three items",
			args{
				elements: []Element{
					{Name: "guitar", Size: 1, Price: 1500},
					{Name: "recorder", Size: 4, Price: 3000},
					{Name: "notebook", Size: 3, Price: 2000},
				},
				totalSize: 4,
			},
			[][]int{
				{1500, 1500, 1500, 1500},
				{1500, 1500, 1500, 3000},
				{1500, 1500, 2000, 3500},
			},
			3500,
		},
		{
			"four items",
			args{
				elements: []Element{
					{Name: "guitar", Size: 1, Price: 1500},
					{Name: "recorder", Size: 4, Price: 3000},
					{Name: "notebook", Size: 3, Price: 2000},
					{Name: "iphone", Size: 1, Price: 2000},
				},
				totalSize: 4,
			},
			[][]int{
				{1500, 1500, 1500, 1500},
				{1500, 1500, 1500, 3000},
				{1500, 1500, 2000, 3500},
				{2000, 3500, 3500, 4000},
			},
			4000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DynamicKnapsack(tt.args.elements, tt.args.totalSize)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantMaxPrice, got.maxPrice())
		})
	}
}
