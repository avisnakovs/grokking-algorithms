package grokking_algorithms

import (
	"math"

	"github.com/pkg/errors"
)

var (
	ErrBadInput = errors.New("bad input")
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
