package main

// 多源最短路算法
import (
	"fmt"
)


const  INFINITY  = 65535

type Graph struct {
	Nv int
	Ne int
	G [][]int
}

type Enode struct {
	V1, V2 int
	Weight int
}

func NewEnode() *Enode {
	return &Enode{}
}

// BuildGraph 邻接矩阵
func BuildGraph(N int) (Mgraph *Graph) {
	Mgraph = &Graph{}
	Mgraph.Nv = N
	// 初始化邻接矩阵
	Mgraph.G = make([][]int, N)
	for V := 0; V < N; V ++ {
		Mgraph.G[V] = make([]int, N)
		for W := 0; W < N; W ++ {
			Mgraph.G[V][W] = INFINITY
		}
	}
	// edge 边
	var M int
	fmt.Scanf("%d",&M)
	Mgraph.Ne = M

	for i := 0; i < M; i ++ {
		Edge := NewEnode()
		fmt.Scanf("%d %d %d", &Edge.V1, &Edge.V2, &Edge.Weight)
		Edge.V1--
		Edge.V2--
		InsertEdge(Mgraph, Edge)
	}
	return Mgraph
}

func InsertEdge(Mgraph *Graph, Edge *Enode ) {
	Mgraph.G[Edge.V1][Edge.V2] = Edge.Weight
	Mgraph.G[Edge.V2][Edge.V1] = Edge.Weight
}


func FindAnimal(Mgraph *Graph) {
	var MaxDist, MinDist, animal int
	D :=  Floyd(Mgraph)

	MinDist = INFINITY
	Nv := Mgraph.Nv
	for i := 0; i < Nv; i ++ {
		MaxDist = FindMax(Nv, D, i)
		if MaxDist == INFINITY {
			fmt.Println("0")
			return
		}
		if MaxDist < MinDist { // 找到最长距离更小的动物
			MinDist = MaxDist
			animal = i + 1
		}
	}
	fmt.Println(animal,MinDist)


}
func Floyd( Graph *Graph) (D [][]int){
	var i, j, k int

	/* 初始化 */
	D = Graph.G

	for k = 0; k < Graph.Nv; k++ {
		for i=0; i < Graph.Nv; i++ {
			for j=0; j<Graph.Nv; j++ {
				if( D[i][k] + D[k][j] < D[i][j] ) {
					D[i][j] = D[i][k] + D[k][j]
				}
			}
		}
	}
	return D

}

func FindMax(Nv int ,D [][]int, i int ) (maxDist int){
	maxDist = 0
	for j := 0; j < Nv; j ++ {
		if i != j && maxDist < D[i][j] {
			maxDist = D[i][j]
		}
	}
	return maxDist
}

func main() {
	// Vertex 点
	var N int
	fmt.Scanf("%d",&N)

	// 建立图
	Mgraph := BuildGraph(N)
	//
	//Find()
	FindAnimal(Mgraph)
}

