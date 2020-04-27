package main

import (
	"fmt"
	"math"
)
// 数组方式实现的 完全二叉树
type HeapStruct struct {
	 ElementType []int  // 存储堆元素的数组
	 Size int  // 堆的当前元素个数
	 Capacity int    // 堆的最大容量
}

// NewHeap 以最大堆为例
func NewHeap(MaxSize int ) *HeapStruct {
	H := &HeapStruct{
		ElementType: make([]int, MaxSize+1),
		Size:        0,
		Capacity:    MaxSize,
	}
	// 定义哨兵为大于堆中所有元素的值, 便于以后更快操作
	H.ElementType[0] = math.MaxInt64
	return H
}

func (H *HeapStruct) IsFull() bool {
	return  H.Size == H.Capacity
}

func (H *HeapStruct) Insert(val int) {
	if H.IsFull() {
		fmt.Println("堆已满")
		return
	}

	H.Size++ // 指向堆的下一个 位置
	i := H.Size // i 为插入后 堆中的最后一个元素的位置

	// 维持有序性
	for ; H.ElementType[i/2] < val; i /= 2 {
		H.ElementType[i] = H.ElementType[i/2] // 向下过滤结点, 画图会更清晰
	}
	H.ElementType[i] = val
}

func (H *HeapStruct)IsEmpty() bool {
	return  H.Size == 0
}

func (H *HeapStruct) Delete() int {
	if H.IsEmpty() {
		fmt.Println("堆已空")
		return -1
	}
	MaxItem := H.ElementType[1]

	// H.szie--  先减和后减的区别 目前不清楚
	temp := H.ElementType[H.Size]
	H.Size--

	var (
		Parent int
		Child int
	)
	// Partent * 2 <= H.Size  如果2i 大于size 那就没有左儿子,自然没有右儿子
	for Parent = 1 ; Parent * 2  <= H.Size; Parent = Child {
		Child = Parent * 2
		// 判断左右儿子, 并选一个大的 Child != Size就是判别有没有右儿子
		if Child != H.Size && H.ElementType[Child] < H.ElementType[Child + 1] {
			Child++
		}

		if temp >= H.ElementType[Child] {
			break
		}else{
			H.ElementType[Parent] = H.ElementType[Child]
		}

		// Parent 就到了左右儿子的 值大的那个位置
	}

	H.ElementType[Parent] = temp
	return  MaxItem
}


func (H *HeapStruct) CreateHeap() {
	// 先建立二叉树
	// 已在main函数完成

	// 从最底部的位置 开始调整堆
	// 从最后一个结点的 父结点开始 到根结点1
	for i := H.Size/2; i > 0; i -- {
		PercDown(H, i)
	}
}

func PercDown(H *HeapStruct, p int) {
	// 下滤, 将H中以 H->Data[p] 为根的子堆 调整为最大堆
	var (
		Parent int
		Child int
	)
	X := H.ElementType[p] // 取出根结点 存放的位置

	for Parent = p;  Parent * 2 <= H.Size; Parent = Child{
		Child = Parent * 2
		if Child != H.Size && H.ElementType[Child] < H.ElementType[Child + 1] {
			Child++
		}
		if X >= H.ElementType[Child] {
			break
		}else{
			H.ElementType[Parent] = H.ElementType[Child]
		}
	}
	H.ElementType[Parent] = X

}

func main() {
	Items := []int{
		42,4354,213,324,1213,656745,32,3,5,12,7,
	}

	heap := NewHeap(11)

	for i := 1; i <= len(Items); i ++ {
		heap.ElementType[i] = Items[i-1]
		heap.Size++
	}

	heap.CreateHeap()
	fmt.Println(heap)
	heap.Delete()
	fmt.Println(heap)
	heap.Delete()
	fmt.Println(heap)
	heap.Delete()
	fmt.Println(heap)
	heap.Delete()
	fmt.Println(heap)

	fmt.Println(heap.ElementType[heap.Size])
}
