package main

import "fmt"

type Graph struct {
	vertices map[string][]string
}

// NewGraph создает новый граф
func NewGraph() *Graph {
	return &Graph{make(map[string][]string)}
}

// AddVertex добавляет вершину в граф
func (g *Graph) AddVertex(v string) {
	if _, exists := g.vertices[v]; !exists {
		g.vertices[v] = []string{}
	}
}

// AddEdge добавляет направленное ребро от вершины v к вершине w
func (g *Graph) AddEdge(v, w string) {
	if _, exists := g.vertices[v]; exists {
		g.vertices[v] = append(g.vertices[v], w)
	}
}

// ShortestPath находит кратчайший путь между начальной вершиной startи конечной вершиной end, использует алгоритм BFS
func (g *Graph) ShortestPath(start, end string) ([]string, bool) {
	// Проверяем, сущесвуют ли начальная и конечная вершина в графе
	if _, exists := g.vertices[start]; !exists {
		return nil, false
	}
	if _, exists := g.vertices[end]; !exists {
		return nil, false
	}

	// Используем очередь для хранения вершин, которые нужно исследовать
	queue := []string{start}
	// Используем map для хранения посещенных вершин и пути к ним
	visited := map[string]string{start: ""}

	for len(queue) > 0 {
		// Достаем вершину из начала очереди
		current := queue[0]
		queue = queue[1:]

		// Если достигли конечной  вершины, восстанавливаем путь
		if current == end {
			path := []string{}
			for at := end; at != ""; at = visited[at] {
				path = append([]string{at}, path...)
			}

			return path, true
		}

		// Обходим соседей текущей вершины
		for _, neighbor := range g.vertices[current] {
			if _, seen := visited[neighbor]; !seen {
				queue = append(queue, neighbor)
				visited[neighbor] = current
			}
		}
	}

	return nil, false
}

// PrintGraph выводит граф
func (g *Graph) PrintGraph() {
	for v, e := range g.vertices {
		fmt.Printf("%s -> %v\n", v, e)
	}
}

func main() {
	graph := NewGraph()
	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddVertex("D")

	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "C")
	graph.AddEdge("C", "D")
	graph.AddEdge("D", "A")

	graph.PrintGraph()

	start, end := "A", "D"
	if path, found := graph.ShortestPath(start, end); found {
		fmt.Printf("Кратчайший путь от %s до %s: %v\n", start, end, path)
	} else {
		fmt.Printf("Путь от %s до %s не найден\n", start, end)
	}
}
