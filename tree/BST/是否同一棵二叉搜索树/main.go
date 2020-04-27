package main

import "fmt"

type TreeNode struct {
	Val int
	Left, Right *TreeNode
	flag bool
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}
func main() {
	/*
	对每次组数据
	- 读入N和L
	- 根据第一行序列建树T
	- 依据树T分别判别后面的L个序列是否能与T形成同一搜索树并输出结果
	 */
	var (
		N, L int
	)
	for {
		fmt.Scanf("%d", &N)
		if N == 0 {
			break
		}
		fmt.Scanf("%d", &L)

		T := BuildTree(N)
		for i := 0; i < L; i ++ {
			if judge(T, N) {
				fmt.Println("Yes")
			}else{
				fmt.Println("No")
			}
			ResetTree(T)
		}
	}
}

//func reversal(T *TreeNode) {
//	if T == nil {
//		return
//	}
//	reversal(T.Left)
//	reversal(T.Right)
//	fmt.Println(T.Val)
//}

/*
需要设计的主要函数
- 读数据 建搜索树T
- 判别一序列是否与T构成一样的搜索树
 */

func BuildTree(N int) *TreeNode {
	var val int
	var T *TreeNode
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &val)
		T = insert(T, val)
	}
	return T
}

func insert(T *TreeNode, val int) *TreeNode  {
	if T == nil {
		T = NewTreeNode(val)
	}else{
		if val > T.Val {
			T.Right = insert(T.Right, val)
		}else if val < T.Val {
			T.Left = insert(T.Left, val)
		}
	}
	return T
}
// check 等于BST的find
func check(T *TreeNode,val int) (exist bool ){
	if T.flag {
		if val > T.Val {
			return check(T.Right, val)
		}else if val < T.Val {
			return check(T.Left, val)
		}

	}else{
		if val == T.Val {
			T.flag = true
			exist = true
			return
		}
	}
	return
}

func judge(T *TreeNode, N int ) bool {
	var V int
	var flag = true // flag 存在的意义是 继续读入输入,不然会发生bug  true 代表一致 false 代表不一致
	fmt.Scanf("%d", &V)
	if V != T.Val {
		flag =  false
	}else{
		T.flag = true
	}
	for i := 1; i < N; i++ {
		fmt.Scanf("%d", &V)
		if flag && !check(T,V) {
			flag = false
		}
	}
	return flag
}

// ResetTree 为同一个树 下一次判断整理环境
func ResetTree (T *TreeNode) {
	if T.Left != nil {
		ResetTree(T.Left)
	}
	if T.Right != nil {
		ResetTree(T.Right)
	}
	T.flag = false
}