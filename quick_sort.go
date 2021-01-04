package grokking_algorithms

import "math/rand"

type split struct {
	less    []int
	pivot   int
	greater []int
}

func (s *split) add(slice []int) {
	for _, value := range slice {
		if value < s.pivot {
			s.less = append(s.less, value)
		} else {
			s.greater = append(s.greater, value)
		}
	}
}

func QuickSort(slice []int) []int {
	size := len(slice)
	if size < 2 {
		return slice
	}
	randomIndex := rand.Intn(size)
	s := split{less: make([]int, 0), pivot: slice[randomIndex], greater: make([]int, 0)}
	s.add(slice[:randomIndex])
	s.add(slice[randomIndex+1:])
	return append(append(QuickSort(s.less), s.pivot), QuickSort(s.greater)...)
}
