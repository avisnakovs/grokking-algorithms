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

func (t Table) maxPrice() int {
	return t[len(t)-1][len(t[0])-1]
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
