package main

import "fmt"

type Graph struct {
	vertices []*Vertex
	memo     map[int]int
}

type Vertex struct {
	key       int
	adjacency []*Vertex
}

func (g *Graph) addVertex(k int) {
	previousLength := len(g.vertices)
	g.memo[k] = previousLength
	g.vertices = append(g.vertices, &Vertex{key: k, adjacency: []*Vertex{}})
}

func (g *Graph) addEdge(from int, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid Edge %d --> %d", from, to)
		fmt.Println(err.Error())
		return
	}
	fromVertex.adjacency = append(fromVertex.adjacency, toVertex)
	toVertex.adjacency = append(toVertex.adjacency, fromVertex)
}
func (g *Graph) getVertex(v int) *Vertex {
	if i, ok := g.memo[v]; ok {
		return g.vertices[i]
	}
	return nil
}
func (g *Graph) printGraph() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex : %d with edges ", v.key)
		for _, v := range v.adjacency {
			fmt.Printf(" %d ", v.key)
		}
	}
}

func main() {
	test := &Graph{memo: make(map[int]int), vertices: []*Vertex{}}
	for i := 0; i < 5; i++ {
		test.addVertex(i)
	}
	test.addEdge(0, 2)
	test.addEdge(1, 2)
	test.addEdge(2, 3)
	test.addEdge(4, 3)
	test.addEdge(4, 2)
	test.printGraph()

}
