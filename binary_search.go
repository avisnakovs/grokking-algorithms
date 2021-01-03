package grokking_algorithms

import (
	"github.com/pkg/errors"
)

var (
	NotFound = errors.New("not found")
)

func BinarySearch(search int, slice []int) (int, error) {
	low, high := 0, len(slice)-1
	for low <= high {
		middle := (low + high) / 2
		guess := slice[middle]
		if guess == search {
			return middle, nil
		} else if guess < search {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}
	return 0, NotFound
}
