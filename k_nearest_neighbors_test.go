package grokking_algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateDistance(t *testing.T) {
	type args struct {
		firstRatings  []int
		secondRatings []int
	}
	tests := []struct {
		name         string
		args         args
		wantDistance float64
		wantErr      error
	}{
		{
			"empty",
			args{firstRatings: nil, secondRatings: nil},
			0.0,
			nil,
		},
		{
			"wrong input",
			args{firstRatings: []int{1}, secondRatings: []int{1, 2}},
			0.0,
			ErrBadInput,
		},
		{
			"same rating",
			args{firstRatings: []int{1}, secondRatings: []int{1}},
			0.0,
			nil,
		},
		{
			"same ratings",
			args{firstRatings: []int{1, 2, 3, 4, 5}, secondRatings: []int{1, 2, 3, 4, 5}},
			0.0,
			nil,
		},
		{
			"same ratings",
			args{firstRatings: []int{3, 4, 4, 1, 4}, secondRatings: []int{4, 3, 5, 1, 5}},
			2.0,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDistance, err := CalculateDistance(tt.args.firstRatings, tt.args.secondRatings)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantDistance, gotDistance)
		})
	}
}

func TestGuessRate(t *testing.T) {
	tests := []struct {
		name             string
		neighbourRatings []int
		wantRate         float64
		wantErr          error
	}{
		{"empty", []int{}, 0, ErrEmptyInput},
		{"one neighbor", []int{1}, 1, nil},
		{"five neighbor", []int{5, 4, 4, 5, 3}, 4.2, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRate, err := GuessRate(tt.neighbourRatings)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantRate, gotRate)
		})
	}
}
