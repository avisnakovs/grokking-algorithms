package grokking_algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBfs_Connected(t *testing.T) {
	t.Run("empty graph", func(t *testing.T) {
		g := Graph{}
		got := Bfs(g).Connected("", "")
		assert.False(t, got)
	})

	t.Run("one element not found", func(t *testing.T) {
		start := &Node{Value: "start"}
		g := Graph{Nodes: map[string]*Node{"start": start}}
		got := Bfs(g).Connected("start", "end")
		assert.False(t, got)
	})

	t.Run("self connected to self", func(t *testing.T) {
		start := &Node{Value: "start"}
		g := Graph{Nodes: map[string]*Node{"start": start}}
		got := Bfs(g).Connected("start", "start")
		assert.True(t, got)
	})

	t.Run("connected", func(t *testing.T) {
		end := &Node{Value: "end"}
		start := &Node{Value: "start", Edges: []*Node{end}}
		g := Graph{Nodes: map[string]*Node{"start": start, "end": end}}
		got := Bfs(g).Connected("start", "end")
		assert.True(t, got)
	})

	t.Run("long connection", func(t *testing.T) {
		four := &Node{Value: "four"}
		three := &Node{Value: "three", Edges: []*Node{four}}
		two := &Node{Value: "two"}
		one := &Node{Value: "one", Edges: []*Node{two, three}}
		g := Graph{Nodes: map[string]*Node{"one": one, "two": two, "three": three, "four": four}}
		bfs := Bfs(g)
		oneAndTwo := bfs.Connected("one", "two")
		assert.True(t, oneAndTwo)
		oneAndThree := bfs.Connected("one", "three")
		assert.True(t, oneAndThree)
		oneAndFour := bfs.Connected("one", "four")
		assert.True(t, oneAndFour)
		twoAndThree := bfs.Connected("two", "three")
		assert.False(t, twoAndThree)
		twoAndFour := bfs.Connected("two", "four")
		assert.False(t, twoAndFour)
		threeAndFour := bfs.Connected("three", "four")
		assert.True(t, threeAndFour)
	})

	t.Run("circular connection", func(t *testing.T) {
		three := &Node{Value: "three"}
		two := &Node{Value: "two", Edges: []*Node{three}}
		one := &Node{Value: "one", Edges: []*Node{two, three}}
		three.Edges = append(three.Edges, one, two)
		two.Edges = append(three.Edges, one)
		g := Graph{Nodes: map[string]*Node{"one": one, "two": two, "three": three}}
		bfs := Bfs(g)
		oneAndTwo := bfs.Connected("one", "two")
		assert.True(t, oneAndTwo)
		oneAndThree := bfs.Connected("one", "three")
		assert.True(t, oneAndThree)
		twoAndThree := bfs.Connected("two", "three")
		assert.True(t, twoAndThree)
	})
}
