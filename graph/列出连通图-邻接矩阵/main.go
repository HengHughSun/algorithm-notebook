package main

import "fmt"

type WeightType int
type Vertex int

type Gnode struct {
	// 顶点数
	Nv Vertex
	// 边数
	Ne Vertex
	G [][]WeightType
}

type Grapha Gnode

func NewGrapha() *Grapha {
	return &Grapha{}
}


// CreateGraph 建立图  初始化一个有VertexNum个顶点但没有边的图
func CreateGraph(Nv Vertex) (Mgrapha *Grapha) {
	Mgrapha = NewGrapha()

	var V Vertex
	Mgrapha.G = make([][]WeightType, Nv)
	for V = 0; V < Nv; V++ {
		Mgrapha.G[V] = make([]WeightType, Nv)
		//for W = 0; W < Mgrapha.Nv; W++ {
		//	Mgrapha.G[V][W]  = 0 // or INFINITY
		//}
	}
	return Mgrapha
}

type ENode struct {
	V1, V2 Vertex   // 有向边<v1, v2>
	//Weight WeightType // 权重
}

type Edge ENode

// InsertEdge 插入边
func InsertEdge(Grapha *Grapha, E *Edge) {
	// 插入边<v1, v2>
	Grapha.G[E.V1][E.V2] = 1

	// 若是无向图, 还要插入边<v2, v1>
	Grapha.G[E.V2][E.V1] = 1
}


func main() {
	// Mgrapha  邻接矩阵
	var Mgrapha * Grapha

	var Nv Vertex
	fmt.Scanf("%d", &Nv)

	Mgrapha = CreateGraph(Nv)
	Mgrapha.Nv = Nv

	fmt.Scanf("%d", &Mgrapha.Ne)
	if Mgrapha.Ne != 0 {
		for i := Vertex(0); i < Mgrapha.Ne; i ++ {
			var V1, V2 Vertex
			fmt.Scanf("%d %d\n",&V1, &V2)
			InsertEdge(Mgrapha, &Edge{V1: V1, V2: V2})

		}
	}

	 DFSVisited = make(map[Vertex]bool)
	 BFSVisited = make(map[Vertex]bool)
	ListComponentsDFS(Mgrapha)
	ListComponentsBFS(Mgrapha)

	//fmt.Println(Mgrapha.G)
}

func ListComponentsDFS(mgrapha *Grapha) {
	for i := Vertex(0); i < mgrapha.Nv; i++{
		if !DFSVisited[i] {
			fmt.Print("{ ")
			DFS(mgrapha, i)
			fmt.Println("}")
		}
	}

}
func ListComponentsBFS(mgrapha *Grapha) {
	for i := Vertex(0); i < mgrapha.Nv; i++{
		if !BFSVisited[i] {
			fmt.Print("{ ")
			BFS(mgrapha, i)
			fmt.Println("}")
		}
	}

}
var DFSVisited map[Vertex]bool
func DFS(grapha *Grapha, v Vertex) {
	DFSVisited[v] = true
	fmt.Printf("%d ", v)
	for w := Vertex(0); w < grapha.Nv ; w ++ {
		if grapha.G[v][w] == 1 && !DFSVisited[w]{
			DFS(grapha, w)
		}
	}
}

var BFSVisited map[Vertex]bool
var c = make(chan Vertex, 300)
func BFS(grapha *Grapha, v Vertex) {
	//Or however many you might need + buffer.
	//c := make(chan int, 300)
	//
	////Push
	//c <- value
	//
	////Pop
	//x <- c
	BFSVisited[v] = true
	c <- v
	fmt.Printf("%d ", v)
	for len(c) != 0 {
	tmp := <- c
		for w := Vertex(0); w < grapha.Nv ; w ++ {
			if grapha.G[tmp][w] == 1 && !BFSVisited[w]{
				fmt.Printf("%d ", w)
				BFSVisited[w] = true
				c <- Vertex(w)
			}
		}
	//default:
	//	return
	}
}
