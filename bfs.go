package grokking_algorithms

import "container/list"

type Graph struct {
	Nodes map[string]*Node
}

type Node struct {
	Value string
	Edges []*Node
}

type Bfs Graph

func (b Bfs) Connected(start string, target string) bool {
	startNode, found := b.Nodes[start]
	if !found {
		return false
	}
	q := newQueue(startNode)
	visitedNodes := make(map[string]struct{})
	for q.len() > 0 {
		node := q.pop()
		if _, visited := visitedNodes[node.Value]; !visited {
			if node.Value == target {
				return true
			}
			visitedNodes[node.Value] = struct{}{}
			q.pushEdges(node.Edges)
		}
	}
	return false
}

func newQueue(startNode *Node) queue {
	l := list.New()
	l.PushBack(startNode)
	return queue{l}
}

type queue struct {
	list *list.List
}

func (q queue) pop() *Node {
	element := q.list.Front()
	q.list.Remove(element)
	return element.Value.(*Node)
}

func (q queue) pushEdges(edges []*Node) {
	for _, edge := range edges {
		q.push(edge)
	}
}

func (q queue) push(edge *Node) {
	q.list.PushBack(edge)
}

func (q queue) len() int {
	return q.list.Len()
}
