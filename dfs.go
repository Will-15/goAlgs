package main

import "fmt"

var visited []bool

func (g *Graph) DFS(start int) {
	visited := make([]bool, g.vertices)
	g.dfsUtil(start, visited)
	fmt.Println()
}

func (g *Graph) dfsUtil(node int, visited []bool) {
	visited[node] = true
	fmt.Print(node, " ")

	for _, neighbor := range g.adjList[node] {
		if !visited[neighbor] {
			g.dfsUtil(neighbor, visited)
		}
	}
}
