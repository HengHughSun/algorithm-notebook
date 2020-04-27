package main

import (
	"fmt"
)

type Graph  struct {
	// 顶点数
	Nv int
	// 边数
	Ne  int
	G []*node
}

func NewGraph() *Graph {
	return &Graph{}
}
type node struct {
	V int // 结点的编号
	Next *node
}
//Newnode ...
func Newnode(V int) *node {
	return &node{
		V:V,
	}
}
// CreateGraph 建立图  初始化一个有VertexNum个顶点但没有边的图
func CreateGraph(Nv int) (Lgraph *Graph) {
	Lgraph = NewGraph()

	Lgraph.G = make([]*node, Nv + 1)
	for V := 1; V < Nv + 1; V++ {
		Lgraph.G[V] = &node{V:V}
	}
	return Lgraph
}

// InsertEdge 插入边
func InsertEdge(Lgraph *Graph, v1, v2 int ) {
	// 插入边<v1, v2>
	tmp := Newnode(v2)
	tmp.Next = Lgraph.G[v1].Next
	Lgraph.G[v1].Next = tmp

	// 若是无向图, 还要插入边<v2, v1>
	tmp = Newnode(v1)
	tmp.Next = Lgraph.G[v2].Next
	Lgraph.G[v2].Next = tmp
}


func main() {
	// Lgraph  邻接表
	var Lgraph * Graph


	var Nv int
	fmt.Scanf("%d", &Nv)

	Lgraph = CreateGraph(Nv)
	Lgraph.Nv = Nv

	fmt.Scanf("%d", &Lgraph.Ne)

	for i := 0; i < Lgraph.Ne; i ++ {
		var x, y int
		fmt.Scanf("%d %d\n",&x, &y)
		InsertEdge(Lgraph, x, y)
	}
	ListComponentsBFS(Lgraph)
}

var c chan int

func ListComponentsBFS(Lgraph *Graph) {
	var count int
	c = make(chan int, 300)
	for i := 1; i < Lgraph.Nv + 1; i++{
			count = BFS(Lgraph, i)
		// 这里是个坑
		//fmt.Printf("%d: %.2f%%\n",i,float64(count*100/Lgraph.Nv))
		fmt.Printf("%d: %.2f%%\n",i,float64(count)*100.0/float64(Lgraph.Nv))
	}
}


func BFS(Lgraph *Graph, i int ) int {
	var Visited map[int]bool
	Visited = make(map[int]bool)

	var count int
	Visited[i] = true

	count = 1
	level := 0
	last := i
	tail := 0

	c <- i // 入队列

	for len(c) > 0 {
		k := <- c
		tmp := Lgraph.G[k]
		for ; tmp != nil; tmp = tmp.Next {
			if !Visited[tmp.V] {
				Visited[tmp.V] = true
				c <- tmp.V
				tail = tmp.V
				count++
			}
		}
		if Lgraph.G[k].V == last {
			level++
			last = tail
		}
		if level == 6 {
			break
		}
	}

	for len(c) > 0 {
		<- c
	}
	return count
}
