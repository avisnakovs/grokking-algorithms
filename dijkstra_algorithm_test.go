package grokking_algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDijkstra_Find(t *testing.T) {
	t.Run("empty graph", func(t *testing.T) {
		wg := WeightedGraph{Nodes: map[string]*WeightedNode{}}

		got, found := Dijkstra(wg).Find("start", "end")

		assert.False(t, found)
		assert.Equal(t, Weight(0), got)
	})

	t.Run("not connected", func(t *testing.T) {
		start := &WeightedNode{Value: "start"}
		a := &WeightedNode{Value: "a"}
		end := &WeightedNode{Value: "end"}

		start.Edges = map[*WeightedNode]Weight{a: 6}
		a.Edges = map[*WeightedNode]Weight{}

		wg := WeightedGraph{Nodes: map[string]*WeightedNode{"start": start, "a": a, "end": end}}

		got, found := Dijkstra(wg).Find("start", "end")
		assert.False(t, found)
		assert.Equal(t, Weight(0), got)
	})

	t.Run("self connected to self", func(t *testing.T) {
		start := &WeightedNode{Value: "start"}

		wg := WeightedGraph{Nodes: map[string]*WeightedNode{"start": start}}

		got, found := Dijkstra(wg).Find("start", "start")

		assert.True(t, found)
		assert.Equal(t, Weight(0), got)
	})

	t.Run("4-nodes graph", func(t *testing.T) {
		start := &WeightedNode{Value: "start"}
		a := &WeightedNode{Value: "a"}
		b := &WeightedNode{Value: "b"}
		end := &WeightedNode{Value: "end"}

		start.Edges = map[*WeightedNode]Weight{a: 6, b: 2}
		a.Edges = map[*WeightedNode]Weight{end: 1}
		b.Edges = map[*WeightedNode]Weight{a: 3, end: 5}

		wg := WeightedGraph{Nodes: map[string]*WeightedNode{"start": start, "a": a, "b": b, "end": end}}

		got, found := Dijkstra(wg).Find("start", "end")
		assert.True(t, found)
		assert.Equal(t, Weight(6), got)
	})

	t.Run("6-node graph", func(t *testing.T) {
		book := &WeightedNode{Value: "book"}
		record := &WeightedNode{Value: "record"}
		poster := &WeightedNode{Value: "poster"}
		guitar := &WeightedNode{Value: "guitar"}
		drum := &WeightedNode{Value: "drum"}
		piano := &WeightedNode{Value: "piano"}

		book.Edges = map[*WeightedNode]Weight{record: 5, poster: 0}
		record.Edges = map[*WeightedNode]Weight{drum: 20, guitar: 15}
		poster.Edges = map[*WeightedNode]Weight{guitar: 30, drum: 35}
		guitar.Edges = map[*WeightedNode]Weight{piano: 20}
		drum.Edges = map[*WeightedNode]Weight{piano: 10}

		wg := WeightedGraph{Nodes: map[string]*WeightedNode{
			"book": book, "record": record, "poster": poster, "drum": drum, "guitar": guitar, "piano": piano}}

		got, found := Dijkstra(wg).Find("book", "piano")
		assert.True(t, found)
		assert.Equal(t, Weight(35), got)
	})

	t.Run("circular connection", func(t *testing.T) {
		start := &WeightedNode{Value: "start"}
		a := &WeightedNode{Value: "a"}
		b := &WeightedNode{Value: "b"}
		end := &WeightedNode{Value: "end"}

		start.Edges = map[*WeightedNode]Weight{a: 6, b: 2}
		a.Edges = map[*WeightedNode]Weight{b: 1, end: 1}
		b.Edges = map[*WeightedNode]Weight{a: 3, end: 5}
		end.Edges = map[*WeightedNode]Weight{a: 3}

		wg := WeightedGraph{Nodes: map[string]*WeightedNode{"start": start, "a": a, "b": b, "end": end}}

		got, found := Dijkstra(wg).Find("start", "end")
		assert.True(t, found)
		assert.Equal(t, Weight(6), got)
	})
}
