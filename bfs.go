package main

import "fmt"

type Graph struct {
	vertices int
	adjList  [][]int
}

// NewGraph creates a new graph with a specified number of vertices.
// It initializes an adjacency list to store edges for each vertex.

func NewGraph(v int) *Graph {
	return &Graph{
		vertices: v,
		adjList:  make([][]int, v),
	}
}

// AddEdge adds a new edge between two vertices in the graph. For an undirected graph,
// it adds edges in both directions.
func (g *Graph) AddEdge(src, dest int) {
	g.adjList[src] = append(g.adjList[src], dest)
	g.adjList[dest] = append(g.adjList[dest], src) // For undirected graph
}

func (g *Graph) BFS(start int) {
	visited := make([]bool, g.vertices)
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Print(node, " ")

		for _, neighbor := range g.adjList[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	fmt.Println()
}
