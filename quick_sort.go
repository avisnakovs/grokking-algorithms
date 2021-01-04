package grokking_algorithms

func QuickSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	pivot := slice[0]
	less, greater := make([]int, 0), make([]int, 0)
	for _, i := range slice[1:] {
		if i < pivot {
			less = append(less, i)
		} else {
			greater = append(greater, i)
		}
	}
	return append(append(QuickSort(less), pivot), QuickSort(greater)...)
}
