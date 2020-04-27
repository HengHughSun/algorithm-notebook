package main

import (
	"fmt"
	"sort"
)


//type heap struct {
//	Data []int
//	MaxSize int
//	Size  int
//}
//
//func NewHeap(MaxSize int) *heap {
//	h := &heap{
//		Data: make([]int,MaxSize + 1),
//		MaxSize: MaxSize,
//		Size: 0,
//	}
//
//	h.Data[0] = math.MinInt8
//	return h
//}
//
//// Insert 模仿 堆的思想
//func (h *heap)Insert(val int) {
//	h.Size++
//	h.Data[h.Size] = val
//
//	var i int
//	 i = h.Size
//	for ; h.Data[i/2] > val; i = i /2  {
//		h.Data[i] = h.Data[i/2]
//	}
//	h.Data[i] = val
//}



var l int
var In []int
var Level []int
var t int = 0
func main() {
	fmt.Scanf("%d", &l)

	In = make([]int, l)

	// 因为是完全二叉树, 所以可以确定 根结点的位置 是l/2
	// 其实先序遍历
	for i := 0; i < l ; i ++ {
		var val int
		fmt.Scanf("%d", &val)
		In[i] = val
	}
	sort.Ints(In)
	Level = make([]int,l)
	solve(0)

	// print
	fmt.Printf("%d", Level[0])
	for i := 1; i < l; i ++ {
		fmt.Printf(" %d", Level[i])
	}

}

// 中序遍历 转 层序遍历
func solve(root int) {
	if root >= l {
		return
	}
	solve(root * 2 + 1)
	// 先到左子树
	Level[root] = In[t]
	t ++
	solve(root * 2 + 2)
}
//func solve(Aleft int, Aright int, Troot int) {
//	n := Aright - Aleft + 1
//	if n <= 0 {return}
//	L := n /2
//	T[Troot] = A[L]
//	LeftTroot := Troot * 2 + 1
//	RightTroot := LeftTroot + 1
//	solve(Aleft, L -1, LeftTroot )
//	solve(L +1, Aright, RightTroot)
//}

