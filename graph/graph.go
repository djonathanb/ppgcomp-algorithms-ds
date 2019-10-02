package graph

import "github.com/djonathanb/ppgcomp-algorithms-ds/ds"

// VerticesContainer defines the contract of vertices containers implementations
type VerticesContainer interface {
	isAdjacent(i int, j int) bool
}

// ListContainer edges linked list implementation
type ListContainer struct {
	edges []ds.LinkedList
}

// NewListContainer creates a new linked list to store edges
func NewListContainer(verticesCount int) ListContainer {
	edges := make([]ds.LinkedList, verticesCount)

	for i := 0; i < verticesCount; i++ {
		edges[i] = ds.NewLinkedList()
	}

	return ListContainer{
		edges: edges,
	}
}

// InsertVertice adds a new vertice to the container
func (l *ListContainer) InsertVertice(i int, j int) {
	l.edges[i].Insert(j)
}

// IsAdjacent verify if two vertices are adjacent
func (l *ListContainer) IsAdjacent(i int, j int) bool {
	edge := l.edges[i].Search(j)
	return edge != nil
}

// MatrixContainer edges matrix implementation
type MatrixContainer struct {
	edges [][]bool
}

// NewMatrixContainer create a new matrix [vertices_count, vertices_count] of size
func NewMatrixContainer(verticesCount int) MatrixContainer {
	matrix := MatrixContainer{
		edges: make([][]bool, verticesCount),
	}

	for i := 0; i < verticesCount; i++ {
		matrix.edges[i] = make([]bool, verticesCount)
	}

	return matrix
}

// NewMatrixContainerFrom returns a new matrix with given edges
func NewMatrixContainerFrom(edges [][]bool) MatrixContainer {
	return MatrixContainer{
		edges: edges,
	}
}

// IsAdjacent verify if two vertices are adjacent
func (l *MatrixContainer) IsAdjacent(i int, j int) bool {
	return l.edges[i][j]
}

// Vertex vertex implementation
type Vertex struct {
	key  int
	data interface{}
}

// NewVertex creates a new vertex
func NewVertex(key int, data interface{}) Vertex {
	return Vertex{
		key:  key,
		data: data,
	}
}

// Graph graph implementation
type Graph struct {
	vertices []Vertex
	edges    VerticesContainer
}

// NewGraph creates a new graph given vertices and edges
func NewGraph(vertices []Vertex, edges VerticesContainer) Graph {
	return Graph{
		vertices: vertices,
		edges:    edges,
	}
}
