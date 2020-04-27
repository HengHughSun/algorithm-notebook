package main 
import "fmt"
// 数组方式实现的 完全二叉树
type HeapStruct struct {
	ElementType []HuffmanNode// 存储堆元素的数组
	Size int  // 堆的当前元素个数
	Capacity int    // 堆的最大容量
}

// NewHeap 以最大堆为例
func NewHeap(MaxSize int ) *HeapStruct {
	H := &HeapStruct{
		ElementType: make([]HuffmanNode, MaxSize+1),
		Size:        0,
		Capacity:    MaxSize,
	}
	// 定义哨兵为大于堆中所有元素的值, 便于以后更快操作
	H.ElementType[0] = NewHuffmanNode(' ', -1)
	return H
j

func (H *HeapStruct) IsFull() bool {
	return  H.Size == H.Capacity
}

func (H *HeapStruct) IsEmpty() bool {
	return  H.Size == 0
}

func (H *HeapStruct) Insert(val HuffmanNode) {
	//if H.IsFull() {
	//	fmt.Println("堆已满")
	//	return
	//}

	H.Size++ // 指向堆的下一个 位置
	i := H.Size // i 为插入后 堆中的最后一个元素的位置

	// 维持有序性
	for ; H.ElementType[i/2].Weight > val.Weight; i /= 2 {
		H.ElementType[i] = H.ElementType[i/2] // 向下过滤结点, 画图会更清晰
	}
	H.ElementType[i] = val
}

func (H *HeapStruct) Delete() *HuffmanNode {
	//if H.IsEmpty() {
	//	fmt.Println("堆已空")
	//	return -1
	//}
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
		// 判断左右儿子, 并选一个小的 Child != Size就是判别有没有右儿子
		if Child != H.Size && H.ElementType[Child].Weight > H.ElementType[Child + 1].Weight {
			Child++
		}

		if temp.Weight <= H.ElementType[Child].Weight {
			break
		}else{
			H.ElementType[Parent] = H.ElementType[Child]
		}

		// Parent 就到了左右儿子的 值大的那个位置
	}

	H.ElementType[Parent] = temp
	return  &MaxItem
}

type HuffmanNode struct {
	Weight int
	Data  byte
	Left, Right *HuffmanNode
}

func NewHuffmanNode(c byte, weight int) HuffmanNode {
	return HuffmanNode{
		Weight: weight,
		Data:c,
	}
}

var (
	wlp   int
	originMap map[byte]int
	N     int

)

func main() {
	// 读入数据
	fmt.Scanf("%d\n", &N)
	H := NewHeap(N)

	// 建一个一个map保存 初始值
	originMap = make(map[byte]int)
	// 建立最小树
	for i := 0 ; i < N; i ++ {
		var (
			c byte
			weight int
		)
		fmt.Scanf("%c %d", &c, &weight)
		originMap[c] = weight
		H.Insert(NewHuffmanNode(c, weight))
	}
	// 最小堆 取出两个最小的 相加称为新的根结点

	//模拟构建huffman树过程：每次从堆中弹出权值最小的两个子树进行合并，合并后的树（权为子树权和）再入堆；
	//而wpl值 = 两个子树的wpl值（权重在子树内的路径和） + 两个子树的权重和（权重又在父子节点间有1的路径长）；
	//因此不需要构建huffman树，只需要保存上述两个数值，其中子树的权重和保存在堆中，wpl变量保存子树的wpl值，同时用于累加
	// 举个例子 有3个结点的huffman tree
	/*
		7	
      3   4
	1  2
	以递归的思想 看最小的一棵树 3 1 2   3就是wpl 因为1 * 1 + 1 * 2 = 3 // 3 又是这个树的最优wpl 所以累加即可
	*/
	for H.Size != 1 {
		node  := NewHuffmanNode(' ', 0)
		node.Left = H.Delete()
		node.Right = H.Delete()
		node.Weight = node.Left.Weight + node.Right.Weight
		wlp += node.Weight  //  666
		H.Insert(node)
	}
	//T := H.Delete()
	// 最终形成huffman树
	check()
}
//func reserval(T *HuffmanNode, count int ) { // 加权不用这样算
//	if T == nil {
//		return
//	}
//	if T.Data != ' ' {
//		sum += T.Weight * count
//	}
//	count++
//	reserval(T.Right, count)
//	reserval(T.Left, count)
//}

func check() {
	var M int
	fmt.Scanf("%d\n", &M)

	for i := 0; i < M; i ++ {
		var c byte
		var res int
		s := make([]string, N)
		for i := 0; i < N; i ++ {
			fmt.Scanf("%c %s",&c, &s[i])
			res += len(s[i]) * originMap[c]
		}
		if wlp == res && !hasPrefixCode(s)  {
			fmt.Println("Yes")
		}else{
			fmt.Println("No")
		}
	}
}

func hasPrefixCode(s []string) bool {
	for i := 0; i < N ; i++ {
		for j := i + 1; j < N; j++ {
			if isPrefix(s[i], s[j]) {
			//if !strings.HasPrefix(s[i], s[j]) && !strings.HasPrefix(s[j], s[i]){
				return true
			}
		}
	}
	return false
}

func isPrefix(s1 string, s2 string) bool {
	var i int
	for i < len(s1) && i < len(s2) && s1[i] == s2[i] {
		i++
	}
	if i == len(s1) || i == len(s2) {
		return true
	}else{
		return false
	}
}
