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
			assert.Equal(t, tt.wantMaxPrice, got.lastValue())
		})
	}
}

func TestDynamicCommonSubstring(t *testing.T) {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		name             string
		args             args
		want             Table
		wantSubstringLen int
	}{
		{
			"fort fosh",
			args{
				first:  "fort",
				second: "fosh",
			},
			[][]int{
				{1, 0, 0, 0},
				{0, 2, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			2,
		},
		{
			"fish fosh",
			args{
				first:  "fish",
				second: "fosh",
			},
			[][]int{
				{1, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 2},
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DynamicCommonSubstring(tt.args.first, tt.args.second)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantSubstringLen, got.maxValue())
		})
	}
}

func TestDynamicCommonSequence(t *testing.T) {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		name            string
		args            args
		want            Table
		wantSequenceLen int
	}{
		{
			"fort fosh",
			args{
				first:  "fort",
				second: "fosh",
			},
			[][]int{
				{1, 1, 1, 1},
				{1, 2, 2, 2},
				{1, 2, 2, 2},
				{1, 2, 2, 2},
			},
			2,
		},
		{
			"fish fosh",
			args{
				first:  "fish",
				second: "fosh",
			},
			[][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 2, 2},
				{1, 1, 2, 3},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DynamicCommonSequence(tt.args.first, tt.args.second)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantSequenceLen, got.lastValue())
		})
	}
}
