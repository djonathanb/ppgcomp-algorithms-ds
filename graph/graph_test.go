package graph

import (
	"testing"

	"github.com/djonathanb/ppgcomp-algorithms-ds/utils"
)

func createMatrixGraph() Graph {
	vertices := []Vertex{
		NewVertex(0, 'A'),
		NewVertex(1, 'B'),
		NewVertex(2, 'C'),
		NewVertex(3, 'D'),
		NewVertex(4, 'E'),
	}

	edges := NewMatrixContainer(len(vertices))
	edges.InsertEdge(0, 1, 1)
	edges.InsertEdge(0, 3, 3)
	edges.InsertEdge(0, 4, 10)
	edges.InsertEdge(1, 2, 5)
	edges.InsertEdge(2, 4, 1)
	edges.InsertEdge(3, 2, 2)
	edges.InsertEdge(3, 4, 6)

	return NewGraph(vertices, edges)
}

func TestGraphMatrixCreation(t *testing.T) {
	graph := createMatrixGraph()
	if &graph == nil || len(graph.vertices) != 5 || graph.edges == nil {
		t.Error("creation failed for matrix graph")
	}
}

func createListGraph() Graph {
	vertices := []Vertex{
		NewVertex(0, 'A'),
		NewVertex(1, 'B'),
		NewVertex(2, 'C'),
		NewVertex(3, 'D'),
		NewVertex(4, 'E'),
	}

	edges := NewListContainer(len(vertices))
	edges.InsertEdge(0, 1, 1)
	edges.InsertEdge(0, 3, 3)
	edges.InsertEdge(0, 4, 10)
	edges.InsertEdge(1, 2, 5)
	edges.InsertEdge(2, 4, 1)
	edges.InsertEdge(3, 2, 2)
	edges.InsertEdge(3, 4, 6)

	return NewGraph(vertices, edges)
}

func TestGraphListCreation(t *testing.T) {
	graph := createListGraph()
	if &graph == nil || len(graph.vertices) != 5 || graph.edges == nil {
		t.Error("creation failed for list graph")
	}
}

func TestGraphListMinimumPathTo(t *testing.T) {
	graph := createListGraph()
	path, weight := graph.MinimumPathTo(0, 4)

	expectedPath := []int{0, 3, 2, 4}
	if utils.Different(path, expectedPath) {
		t.Errorf("path different from %v", expectedPath)
	}

	expectedWeight := 6
	if weight != expectedWeight {
		t.Errorf("weight different from %d", expectedWeight)
	}
}
