package main

import (
	"fmt"
)

type Graph struct {
	vertices map[int][]int
}

func NewGraph() *Graph {
	return  &Graph{
	  vertices:	make(map[int][]int),
	}
}

func (g *Graph) AddEdge(v1,v2 int){
	g.vertices[v1]=append(g.vertices[v1],v2)
	g.vertices[v2]=append(g.vertices[v2],v1)
}

func (g*Graph) RemoveEdge(v1,v2 int){
	g.vertices[v1]=removeFromSlice(g.vertices[v1],v2)
	g.vertices[v2]=removeFromSlice(g.vertices[v2],v1)
}

func removeFromSlice (slice []int,val int) []int{
	for i ,v:=range slice{
		if(v==val){
			return append(slice[:i],slice[i+1:]...)
		}
	}
	return slice
}

func (g* Graph) Print(){
	for ver ,edge:=range g.vertices{
		fmt.Printf("%d: %v\n",ver,edge)
	}
}

func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)  
	queue := []int{start}           

	fmt.Printf("BFS starting from vertex %d: ", start)

	for len(queue) > 0 {
		 
		vertex := queue[0]
		queue = queue[1:]

		 
		if !visited[vertex] {
			 
			visited[vertex] = true
			fmt.Printf("%d ", vertex)

		 
			for _, neighbor := range g.vertices[vertex] {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
				}
			}
		}
	}
	fmt.Println()
}

func (g *Graph) DFS(start int){

	visited := make(map[int]bool)  
	stack := []int{start}           

	fmt.Printf("BFS starting from vertex %d: ", start)

	for len(stack) > 0 {
		 
		vertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		 
		if !visited[vertex] {
			 
			visited[vertex] = true
			fmt.Printf("%d ", vertex)

		 
			for _, neighbor := range g.vertices[vertex] {
				if !visited[neighbor] {
					stack = append(stack, neighbor)
				}
			}
		}
	}
	fmt.Println()
}
func main(){
	graph:=NewGraph()
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
	fmt.Println("graph adj list")
	graph.Print()
	graph.BFS(1)
	graph.DFS(1)
}
