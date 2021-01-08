package grokking_algorithms

import (
	"math"
)

type WeightedGraph struct {
	Nodes map[string]*WeightedNode
}

type WeightedNode struct {
	Value string
	Edges map[*WeightedNode]Weight
}

type Weight int

type Dijkstra WeightedGraph

func (d Dijkstra) Find(start string, target string) (total Weight, found bool) {
	startNode, targetNode, found := d.findNodes(start, target)
	if !found {
		return 0, false
	}
	if startNode.Value == targetNode.Value {
		return 0, true
	}
	path := findPath(startNode, targetNode)
	return path.weight(targetNode)
}

func (d Dijkstra) findNodes(start string, target string) (*WeightedNode, *WeightedNode, bool) {
	startNode, found := d.Nodes[start]
	if !found {
		return nil, nil, false
	}
	targetNode, found := d.Nodes[target]
	if !found {
		return nil, nil, false
	}
	return startNode, targetNode, true
}

func findPath(startNode *WeightedNode, targetNode *WeightedNode) path {
	path := newPath(startNode, targetNode)
	node := path.findSmallestWeight()
	for node != nil {
		path.processNode(node)
		node = path.findSmallestWeight()
	}
	return path
}

func newPath(start *WeightedNode, target *WeightedNode) path {
	path := path{
		processedNodes: make(map[*WeightedNode]struct{}),
		parentNodes:    make(map[*WeightedNode]*WeightedNode),
		nodeWeights:    make(map[*WeightedNode]Weight),
	}
	path.addStartNode(start)
	path.addTargetNode(target)
	return path
}

type path struct {
	processedNodes map[*WeightedNode]struct{}
	parentNodes    map[*WeightedNode]*WeightedNode
	nodeWeights    map[*WeightedNode]Weight
}

func (p path) addStartNode(startNode *WeightedNode) {
	for node, weight := range startNode.Edges {
		p.nodeWeights[node] = weight
		p.parentNodes[node] = startNode
	}
}

func (p path) addTargetNode(targetNode *WeightedNode) {
	p.nodeWeights[targetNode] = math.MaxInt64
}

func (p path) findSmallestWeight() *WeightedNode {
	var result *WeightedNode
	var minWeight Weight = math.MaxInt64
	for node, weight := range p.nodeWeights {
		_, processed := p.processedNodes[node]
		if weight < minWeight && !processed {
			minWeight = weight
			result = node
		}
	}
	return result
}

func (p path) processNode(processedNode *WeightedNode) {
	processedNodeWeight := p.nodeWeights[processedNode]
	for node, edgeWeight := range processedNode.Edges {
		newWeight := processedNodeWeight + edgeWeight
		oldWeight, found := p.nodeWeights[node]
		if !found || newWeight < oldWeight {
			p.nodeWeights[node] = newWeight
			p.parentNodes[node] = processedNode
		}
	}
	p.processedNodes[processedNode] = struct{}{}
}

func (p path) weight(node *WeightedNode) (total Weight, found bool) {
	cost, found := p.nodeWeights[node]
	if !found {
		return 0, false
	}
	if _, hasParent := p.parentNodes[node]; !hasParent {
		return 0, false
	}
	return cost, true
}
