### Graph Implementation with adjacency list

`BFS`, `DFS`, & `Number of Connected Components`

```Go
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
func removeAtIndex[T any](slice []T, s int) []T {
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

func (g *Graph) printEdges() {
	for key, node := range g.vertexMap {
		fmt.Printf("%d ---> %v\n", key, node)
	}
}

func (g *Graph) BFS(v int) {
	visited := map[int]bool{}
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
					fmt.Printf("%d ", val.key)
					queue.PushBack(val.key)
				}
			}
		}
	}
	fmt.Println()
}

func (g *Graph) DFS(v int) {
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
					fmt.Printf("%d ", val.key)
					stack.PushBack(val.key)
				}
			}
		}
	}
	fmt.Println()
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
	graph.printEdges()
	graph.BFS(1)
	graph.DFS(1)

}
//Output

9 ---> [7 8 1]
2 ---> [1 3 4 5 10]
3 ---> [2 4]
5 ---> [1 2]
6 ---> [4]
10 ---> [2]
1 ---> [2 5 9]
4 ---> [2 3 6]
7 ---> [9]
8 ---> [9]
2 5 9 1 3 4 10 7 8 6 
2 5 9 7 8 1 3 4 10 6
```

### Number of Islands
Leetcode URL [https://leetcode.com/problems/number-of-islands/description/](https://leetcode.com/problems/number-of-islands/description/)

```Go
package main

import (
	"container/list"
	"fmt"
)

type Tuple struct {
	x, y int
}

func isLand(grid [][]byte, i, j int) bool {
	return i >= 0 && j >= 0 && j < len(grid[0]) && i < len(grid) && grid[i][j] == 1
}
func numIslands(grid [][]byte) int {
	m, n := len(grid), 0
	if m > 0 {
		n = len(grid[0])
	}
	visited := make([][]bool, m)
	for k := 0; k < m; k++ {
		visited[k] = make([]bool, n)
	}
	q := list.New()
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				q.PushBack(Tuple{i, j})
				for q.Len() > 0 {
					e := q.Front()
					q.Remove(e)
					tuple := e.Value.(Tuple)
					if !visited[tuple.x][tuple.y] {
						visited[tuple.x][tuple.y] = true
						//Top
						if isLand(grid, tuple.x-1, tuple.y) {
							q.PushBack(Tuple{tuple.x - 1, tuple.y})
						}
						//Bottom
						if isLand(grid, tuple.x+1, tuple.y) {
							q.PushBack(Tuple{tuple.x + 1, tuple.y})
						}
						//Right
						if isLand(grid, tuple.x, tuple.y+1) {
							q.PushBack(Tuple{tuple.x, tuple.y + 1})
						}
						//Left
						if isLand(grid, tuple.x, tuple.y-1) {
							q.PushBack(Tuple{tuple.x, tuple.y - 1})
						}
					}
				}
				count++
			}

		}
	}
	return count
}

func isColorChanged(image [][]int, i , j, color int)bool {
    return i >= 0 && j >= 0 && i <len(image) && j < len(image[0]) && image[i][j] == color
}
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    q := list.New()
    q.PushBack(Tuple{sr, sc})
    for q.Len() > 0 {
        e := q.Front()
        t := e.Value.(Tuple)
        q.Remove(e)
        if isColorChanged(image,t.x,t.y,color) {
            image[t.x][t.y] = color
            //Bottom
            if !isColorChanged(image, t.x+1, t.y, color){
                q.PushBack(Tuple{t.x+1, t.y})
            }
            //Top
            if !isColorChanged(image, t.x-1, t.y, color){
                q.PushBack(Tuple{t.x-1, t.y})
            }
            //Right
            if !isColorChanged(image, t.x, t.y+1, color){
                q.PushBack(Tuple{t.x, t.y+1})
            }
            //Left
            if !isColorChanged(image, t.x, t.y-1, color){
                q.PushBack(Tuple{t.x, t.y-1})
            }
        }
    }
    return image
}
func main() {
	// grid := [][]byte{
	// 	{1, 1, 1, 1, 0},
	// 	{1, 1, 0, 1, 0},
	// 	{1, 1, 0, 0, 0},
	// 	{1, 1, 1, 1, 0},
	// }
	// fmt.Println(grid)
	// fmt.Println(numIslands(grid))

	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 1, 0},
		{1, 0, 1},
	}
	floodFill(image, 1,1,2)
	
}
```

### Rotten Oranges
Leetcode URL :[https://leetcode.com/problems/rotting-oranges/](https://leetcode.com/problems/rotting-oranges/)
```Go

package main

import "fmt"

type Tuple struct {
	x, y, time int
}

func isInBound(grid [][]int, i, j int) bool {
	return i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	q := []Tuple{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				visited[i][j] = 2
				q = append(q, Tuple{i, j, 0})
			} else {
				// this else condition is not required as we define int slice which has default as 0
				visited[i][j] = 0
			}
		}
	}
	maxTime := 0
	drow := []int{-1, 0, 1, 0}
	dcol := []int{0, +1, 0, -1}
	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		maxTime = max(maxTime, t.time)
		for i := 0; i < len(drow); i++ {
			nRow := t.x + drow[i]
			nCol := t.y + dcol[i]
			if isInBound(grid, nRow, nCol) && grid[nRow][nCol] == 1 && visited[nRow][nCol] == 0 {
				q = append(q, Tuple{nRow, nCol, t.time + 1})
				visited[nRow][nCol] = 2
			}
		}

	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[i][j] != 2 && grid[i][j] == 1 {
				return -1
			}
		}
	}
	return maxTime
}

func main() {
	// M, N := 3, 3
	// grid := make([][]int, M)
	// for i := 0; i < M; i++ {
	// 	grid[i] = make([]int, N)
	// }
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	fmt.Println(orangesRotting(grid))
}
```