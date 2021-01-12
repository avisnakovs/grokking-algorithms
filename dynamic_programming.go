package grokking_algorithms

type Element struct {
	Name  string
	Size  int
	Price int
}

func NewTable(rows int, columns int) Table {
	result := make([][]int, rows, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]int, columns, columns)
	}
	return result
}

type Table [][]int

func (t Table) getPrice(row, column int) int {
	if row >= 0 && row < len(t) && column >= 0 && column < len(t[0]) {
		return t[row][column]
	}
	return 0
}

func (t Table) setPrice(row, column, price int) {
	if row >= 0 && row < len(t) && column >= 0 && column < len(t[0]) {
		t[row][column] = price
	}
}

func (t Table) lastValue() int {
	return t[len(t)-1][len(t[0])-1]
}

func (t Table) maxValue() int {
	var max int
	for row := 0; row < len(t); row++ {
		for column := 0; column < len(t[0]); column++ {
			if t[row][column] > max {
				max = t[row][column]
			}
		}
	}
	return max
}

func DynamicKnapsack(elements []Element, knapsackSize int) Table {
	table := NewTable(len(elements), knapsackSize)
	for row := 0; row < len(elements); row++ {
		element := elements[row]
		for column := 0; column < knapsackSize; column++ {
			currentSize := column + 1
			previousMax := table.getPrice(row-1, column)
			current := elementPrice(element, currentSize) + table.getPrice(row-1, column-element.Size)
			table.setPrice(row, column, max(previousMax, current))
		}
	}
	return table
}

func elementPrice(element Element, size int) int {
	if size-element.Size < 0 {
		return 0
	}
	return element.Price
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func DynamicCommonSubstring(first, second string) Table {
	table := NewTable(len(first), len(second))
	for row := 0; row < len(first); row++ {
		for column := 0; column < len(second); column++ {
			if first[row] == second[column] {
				previousSubstringLen := table.getPrice(row-1, column-1)
				table.setPrice(row, column, previousSubstringLen+1)
			}
		}
	}
	return table
}

func DynamicCommonSequence(first, second string) Table {
	table := NewTable(len(first), len(second))
	for row := 0; row < len(first); row++ {
		for column := 0; column < len(second); column++ {
			if first[row] == second[column] {
				maxPreviousSequence := table.getPrice(row-1, column-1)
				table.setPrice(row, column, maxPreviousSequence+1)
			} else {
				previousColumn := table.getPrice(row, column-1)
				previousRow := table.getPrice(row-1, column)
				maxPreviousSequence := max(previousRow, previousColumn)
				table.setPrice(row, column, maxPreviousSequence)
			}
		}
	}
	return table
}
