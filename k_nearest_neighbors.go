package grokking_algorithms

import (
	"math"

	"github.com/pkg/errors"
)

var (
	ErrBadInput   = errors.New("bad input")
	ErrEmptyInput = errors.New("empty input")
)

func CalculateDistance(firstRatings, secondRatings []int) (distance float64, err error) {
	if len(firstRatings) != len(secondRatings) {
		return 0, ErrBadInput
	}
	var sum int
	for i := 0; i < len(firstRatings); i++ {
		x := firstRatings[i] - secondRatings[i]
		sum += x * x
	}
	return math.Sqrt(float64(sum)), nil
}

func GuessRate(neighbourRatings []int) (rate float64, err error) {
	if len(neighbourRatings) == 0 {
		return 0, ErrEmptyInput
	}
	var sum int
	for _, rating := range neighbourRatings {
		sum += rating
	}
	return float64(sum) / float64(len(neighbourRatings)), nil
}
