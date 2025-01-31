package main

import (
	"fmt"
)

// Graph represents a graph using an adjacency list
type Graph struct {
	vertices map[int][]int
}

// NewGraph creates a new graph
func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[int][]int),
	}
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(v1, v2 int) {
	g.vertices[v1] = append(g.vertices[v1], v2)
	g.vertices[v2] = append(g.vertices[v2], v1)
}

// RemoveEdge removes an edge from the graph
func (g *Graph) RemoveEdge(v1, v2 int) {
	g.vertices[v1] = removeFromSlice(g.vertices[v1], v2)
	g.vertices[v2] = removeFromSlice(g.vertices[v2], v1)
}

// removeFromSlice removes an element from a slice
func removeFromSlice(slice []int, val int) []int {
	for i, v := range slice {
		if v == val {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// Print prints the adjacency list of the graph
func (g *Graph) Print() {
	for ver, edge := range g.vertices {
		fmt.Printf("%d: %v\n", ver, edge)
	}
}

// BFS performs a breadth-first search starting from the given vertex


func main() {
	graph := NewGraph()
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 3)
	graph.AddEdge(2, 5)
	graph.AddEdge(3, 6)
	graph.AddEdge(4, 5)
	graph.AddEdge(5, 6)
	graph.AddEdge(6, 7)
	graph.AddEdge(7, 8)
	graph.AddEdge(8, 1)
	graph.AddEdge(4, 2)
	graph.AddEdge(3, 5)

	fmt.Println("Graph adjacency list:")
	graph.Print()

	// Perform BFS starting from vertex 1
	graph.BFS(1)
}