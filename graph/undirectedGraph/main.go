package main

import (
	"container/list"
	"fmt"
)

type Graph struct {
	vertices map[string]bool
	edges    map[string]map[string]bool
}

// NewGraph создает новый граф
func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[string]bool),
		edges:    make(map[string]map[string]bool),
	}
}

// AddVertex добавляет вершину в граф
func (g *Graph) AddVertex(v string) {
	if !g.vertices[v] {
		g.vertices[v] = true
		g.edges[v] = make(map[string]bool)
	}
}

// AddEdge добавляет ненаправленные ребра между двумя вершинами
func (g *Graph) AddEdge(v, w string) {
	if g.vertices[v] && g.vertices[w] {
		g.edges[v][w] = true
		g.edges[w][v] = true
	}
}

// RemoveEdge удаляет ненаправленное ребро между двумя вершинами
func (g *Graph) RemoveEdge(v, w string) {
	if g.edges[v] != nil && g.edges[w] != nil {
		delete(g.edges[v], w)
		delete(g.edges[w], v)
	}
}

// BFC выполняет поиск в ширину из заданной начальной вершины
func (g *Graph) BFC(start string) []string {
	if !g.vertices[start] {
		return nil
	}

	visited := make(map[string]bool)
	queue := list.New()
	queue.PushBack(start)
	visited[start] = true

	var result []string
	for queue.Len() > 0 {
		element := queue.Front()
		vertex := element.Value.(string)
		queue.Remove(element)
		result = append(result, vertex)

		for neighbor := range g.edges[vertex] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.PushBack(neighbor)
			}
		}
	}

	return result
}

// Print печатает граф
func (g *Graph) Print() {
	for vertex, edges := range g.edges {
		fmt.Printf("%s:\n", vertex)
		for edge := range edges {
			fmt.Printf(" %s\n", edge)
		}
		fmt.Println()
	}
}

func main() {
	graph := NewGraph()
	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	fmt.Println("Graph:")
	graph.Print()

	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "C")

	fmt.Println("Graph:")
	graph.Print()

	startVertex := "A"
	fmt.Printf("\nBFC starting from %s\n", startVertex)
	bfsResult := graph.BFC(startVertex)
	for _, vertex := range bfsResult {
		fmt.Printf("%s\n", vertex)
	}
	fmt.Println()
}
