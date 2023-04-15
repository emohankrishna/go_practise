package main

import (
	"container/list"
	"fmt"
)

type Graph struct {
	vertexMap map[int][]*Node
}

func NewGraph() *Graph {
	return &Graph{
		vertexMap: make(map[int][]*Node),
	}
}

type Node struct {
	key    int
	weight int
}

func (n Node) String() string {
	return fmt.Sprintf("%d", n.key)
}

func (g *Graph) addVertex(v int) {
	g.vertexMap[v] = []*Node{}
}
func getNode(k int) *Node {
	return &Node{
		key: k,
	}
}

func (g *Graph) isEdgeExist(s, d int) bool {
	list, ok := g.vertexMap[s]
	if ok {
		for _, v := range list {
			if v.key == d {
				return true
			}
		}
	}

	return false
}

func (g *Graph) removeVertex(v int) {
	list := g.vertexMap[v]
	for _, node := range list {
		l := g.vertexMap[node.key]
		if index := getIndex(l, v); index != -1 {
			g.vertexMap[node.key] = removeAtIndex(l, index)
		}
	}
	delete(g.vertexMap, v)
}

func (g *Graph) addEdge(s, d int) {
	sNode := getNode(s)
	dNode := getNode(d)
	if !g.isEdgeExist(s, d) {
		g.vertexMap[s] = append(g.vertexMap[s], dNode)
	}
	if !g.isEdgeExist(d, s) {
		g.vertexMap[d] = append(g.vertexMap[d], sNode)
	}

}

func getIndex(list []*Node, x int) int {
	for i, node := range list {
		if node.key == x {
			return i
		}
	}
	return -1
}
func removeAtIndex(slice []*Node, s int) []*Node {
	return append(slice[:s], slice[s+1:]...)
}

func (g *Graph) removeEdge(s, d int) {
	slist := g.vertexMap[s]
	dlist := g.vertexMap[d]
	if i := getIndex(slist, d); i != -1 {
		g.vertexMap[s] = removeAtIndex(slist, i)
	}
	if i := getIndex(dlist, s); i != -1 {
		g.vertexMap[d] = removeAtIndex(dlist, i)
	}
}

func (g *Graph) printGraph() {
	for key, node := range g.vertexMap {
		fmt.Printf("%d --> %v\n", key, node)
	}
}

func (g *Graph) ConnectedComponents() int {
	visited := make(map[int]bool, len(g.vertexMap))
	q := list.New()
	count := 0
	for k, _ := range g.vertexMap {
		if !visited[k] {
			q.PushBack(k)
			visited[k] = true
			for q.Len() > 0 {
				e := q.Front()
				q.Remove(e)
				adjList := g.vertexMap[e.Value.(int)]
				for _, n := range adjList {
					if !visited[n.key] {
						visited[n.key] = true
						q.PushBack(n.key)
					}
				}

			}
			count++
		}
	}
	return count
}

func (g *Graph) BFS(v int) []int {
	visited := map[int]bool{}
	bfs := []int{}
	queue := list.New()
	queue.PushBack(v)
	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)
		vList, ok := g.vertexMap[element.Value.(int)]
		if ok {
			for _, val := range vList {
				if !visited[val.key] {
					visited[val.key] = true
					bfs = append(bfs, val.key)
					queue.PushBack(val.key)
				}
			}
		}
	}
	return bfs
}

func (g *Graph) DFS(v int) []int {
	dfs := []int{}
	visited := map[int]bool{}
	stack := list.New()
	stack.PushBack(v)
	for stack.Len() > 0 {
		element := stack.Back()
		stack.Remove(element)
		vList, ok := g.vertexMap[element.Value.(int)]
		if ok {
			for _, val := range vList {
				if !visited[val.key] {
					visited[val.key] = true
					dfs = append(dfs, val.key)
					stack.PushBack(val.key)
				}
			}
		}
	}
	return dfs
}
func main() {

	graph := NewGraph()
	graph.addVertex(1)
	graph.addVertex(2)
	graph.addEdge(1, 2)
	graph.addEdge(2, 3)
	graph.addEdge(2, 4)
	graph.addEdge(3, 4)
	graph.addEdge(5, 1)
	graph.addEdge(2, 5)
	graph.addEdge(6, 4)
	graph.addEdge(7, 9)
	graph.addEdge(8, 9)
	graph.addEdge(2, 10)
	graph.addEdge(1, 9)
	graph.addVertex(11)
	graph.addVertex(12)
	graph.printGraph()
	graph.BFS(1)
	graph.DFS(1)
	graph.ConnectedComponents()

}
