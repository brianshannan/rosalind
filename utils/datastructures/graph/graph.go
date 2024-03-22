package graph

import (
	"github.com/brianshannan/rosalind/utils/datastructures/queue"
)

type Graph struct {
	matrix [][]int
}

func NewGraph(numVertices int) *Graph {
	matrix := make([][]int, numVertices)
	for i := 0; i < numVertices; i++ {
		matrix[i] = make([]int, numVertices)
	}
	return &Graph{
		matrix: matrix,
	}
}

func (g *Graph) AddEdge(source, destination int) {
	g.matrix[source][destination] = 1
}

func (g *Graph) BFS(startVertex int, onVertex func(int)) {
	searchQueue := queue.NewQueue[int]()
	searchQueue.Enqueue(startVertex)
	seenVertices := make([]bool, len(g.matrix))

	for !searchQueue.IsEmpty() {
		vertex := searchQueue.Dequeue()
		onVertex(vertex)
		seenVertices[vertex] = true

		for destinationVertex, hasEdge := range g.matrix[vertex] {
			if hasEdge == 0 {
				continue
			}

			if !seenVertices[destinationVertex] {
				searchQueue.Enqueue(destinationVertex)
			}
		}
	}
}

func (g *Graph) FindNumIslands() int {
	unvisitedVertices := make([]bool, len(g.matrix))

	numIslands := 0
	for i := 0; i < len(g.matrix); i++ {
		if unvisitedVertices[i] {
			continue
		}

		numIslands++
		g.BFS(i, func(i int) {
			unvisitedVertices[i] = true
		})
	}
	return numIslands
}
