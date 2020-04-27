package main

// 多源最短路算法
import (
	"fmt"
)

const  INFINITY  = 65535

type node struct {
	X int
	Y int
}

type Grapha  struct {
	// 顶点数
	Nv int
	// 边数
	Ne  int
	G []*node
}

var (
	Mgrapha *Grapha
	//var Visited map[int]bool
	// hypotenuse 斜边
	hypotenuse float64
	Queue      chan int
	Dist       []int
	Path       []int
	first      float64
)

func main() {
	// 如果有多条最短路，选择第一跳最近的。因此先将所有鳄鱼根据离中心点（0，0）的距离来从小到大排序
	// 最小跳数。 所以是无权图

	// 鳄鱼数 the number of crocodiles
	var N int
	fmt.Scanf("%d",&N)

	Mgrapha = &Grapha{}
	Mgrapha.G = make([]*node, N+1)
	Mgrapha.Nv = N

	fmt.Scanf("%d", &Mgrapha.Ne)

	if float64(Mgrapha.Ne) >= 42.5 {
		fmt.Println("1")
		return
	}

	hypotenuse = square(Mgrapha.Ne)

	Mgrapha.G[0] = &node{0,0}
	// 无需建立图 扫描当前结点可以跳到的鳄鱼 BFS   因为DFS找最短路径不是特别合适
	for i := 1; i < N+1; i ++ {
		var x, y int
		fmt.Scanf("%d %d\n",&x, &y)
		Mgrapha.G[i] = &node{x, y}
	}

	Queue = make(chan int, 3000)
	Dist = make([]int, N + 1)
	for i := 0; i < N + 1; i ++{
		Dist[i] = INFINITY
	}

	Path = make([]int, N + 1)
	for i := 0; i < N + 1; i ++{
		Path[i] = -1
	}
	first = square(float64(Mgrapha.Ne) + 7.5)
	ListComponentsBFS()
}

func ListComponentsBFS() {
	var out int // 最后一头鳄鱼的下标
	var short int
	short = INFINITY
	out = 0
	var firstShort float64
	firstShort = INFINITY
	for i := 1; i < Mgrapha.Nv + 1; i ++ {
		if firstJump(i) {
			Queue <- i
			Dist[i] = 1
			Path[i] = 0
			res := Unweighted(i) // 最后一头鳄鱼
			if res > 0 { // 无法到岸的时候 res 返回的是0, 会造成firstshot的更新
				if Dist[res] <= short &&  temp < firstShort {
					short = Dist[res]
					out = res
					firstShort = temp // firstshort 第一跳最短,只能在成功时更新
				}

			}
		}
	}

	if out == 0 {
		fmt.Println("0")
	}else{
		fmt.Println(short + 1)
		help(out)
	}

}
var temp float64

func firstJump(i int) bool {

	temp = square(abs(Mgrapha.G[i].X)) + square(abs(Mgrapha.G[i].Y))
	//if temp > 49 { // 岸上的鳄鱼不算
	if temp > 49 && temp <= first {
			return true
	}

	return  false
}

func help(out int ) {
	tmp := make([]int, 0)

	for Path[out] != -1 {
		tmp = append(tmp, out)
		out = Path[out]
	}

	for i := len(tmp) - 1; i >= 0; i -- {
		fmt.Println(Mgrapha.G[tmp[i]].X, Mgrapha.G[tmp[i]].Y)
	}
}

func Unweighted(i int) (out int) {
	// for V.每一个W {
	for len(Queue) != 0 {
		V := <- Queue // node 的下标
		if IsSave(V) {
			out = V
			break
		}
		for W := 1; W < Mgrapha.Nv + 1; W ++ {
			if V != W && Jump(V, W) {
				if Dist[V] + 1 < Dist[W] {
					Dist[W] = Dist[V] + 1
					Path[W] = V
					Queue <- W
				}
			}
		}
	}

	for len(Queue) != 0 {
		<- Queue // node 的下标
	}
	return out
}

func Jump(v, w int) bool {
	if square(abs(Mgrapha.G[v].X - Mgrapha.G[w].X))	+ square(abs(Mgrapha.G[v].Y - Mgrapha.G[w].Y)) <= hypotenuse {
		return true
	}
	return false
}

func IsSave(v int) bool {
	if Mgrapha.G[v].X + Mgrapha.Ne >= 50 || Mgrapha.G[v].X - Mgrapha.Ne <= -50 || Mgrapha.G[v].Y + Mgrapha.Ne >= 50  || Mgrapha.G[v].Y - Mgrapha.Ne <= -50 {
		return true
	}
	return false
}



func square(n interface{}) float64 {
	switch v := n.(type) {
	case int:
		return float64(v) * float64(v)
	case float64:
		return v * v
	default:
		fmt.Println("invalid var")
	}
	return 0
}

func abs(n int) int {
	if n > 0 {
		return  n
	}
	return -n
}



