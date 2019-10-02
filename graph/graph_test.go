package graph

import "testing"

func TestGraphMatrixCreation(t *testing.T) {
	vertices := []Vertex{
		NewVertex(1, '1'),
		NewVertex(2, '2'),
		NewVertex(3, '3'),
		NewVertex(4, '4'),
		NewVertex(5, '5'),
	}

	edges := NewMatrixContainerFrom([][]bool{
		{false, true, false, false, true},
		{true, false, true, true, true},
		{false, true, false, true, false},
		{false, true, true, false, true},
		{true, true, false, true, false},
	})

	NewGraph(vertices, edges)
}

func TestGraphListCreation(t *testing.T) {
	vertices := []Vertex{
		NewVertex(1, '1'),
		NewVertex(2, '2'),
		NewVertex(3, '3'),
		NewVertex(4, '4'),
		NewVertex(5, '5'),
	}

	edges := NewListContainer(len(vertices))
	edges.InsertVertice(0, 1)
	edges.InsertVertice(0, 4)
	edges.InsertVertice(1, 0)
	edges.InsertVertice(1, 2)
	edges.InsertVertice(1, 3)
	edges.InsertVertice(1, 4)
	edges.InsertVertice(2, 1)
	edges.InsertVertice(2, 3)
	edges.InsertVertice(3, 1)
	edges.InsertVertice(3, 2)
	edges.InsertVertice(3, 4)
	edges.InsertVertice(4, 0)
	edges.InsertVertice(4, 1)
	edges.InsertVertice(4, 3)

	NewGraph(vertices, edges)
}
