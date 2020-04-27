package main

import "fmt"

type TreeNode struct {
	Val int
	Left, Right *TreeNode
	Height int // 树高
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

func Max(a, b int) int{
	if a > b {
		return a
	}
	return b
}

func main() {
	var (
		N int
		val int
		T *TreeNode
	)
	fmt.Scanf("%d", &N)
	for i := 0; i < N; i ++ {
		fmt.Scanf("%d", &val)
		T = CreateTree(T, val)
	}
	fmt.Println(T.Val)
}

// CreateTree ....
func CreateTree(T *TreeNode, val int) *TreeNode {
	return Insert(T, val)
}

// Insert 将X插入AVL树T中, 并且返回调整后的AVL树
func Insert(T *TreeNode, val int) *TreeNode {
	if T == nil { // 空树,新建一个结点
		T = NewTreeNode(val)
	}else if val < T.Val {
		T.Left = Insert(T.Left, val)
		// 如果需要左旋
		if GetHeight(T.Left) - GetHeight(T.Right) == 2 {
			if val < T.Left.Val {
				T = SingleLeftRotation(T)  // 左单旋
			}else{
				T = DoubleLeftRightRotation(T)  // 左-右单旋
			}
		}
	}else if val > T.Val {
		T.Right = Insert(T.Right, val)
		// 如果需要右旋
		if GetHeight(T.Left) - GetHeight(T.Right) == -2 {
			// 被破坏的结点T一定在T.Right的上方
			if val > T.Right.Val { // 破坏者在T.Right的Right 否则左边
				T = SingleRightRotation(T)  // 右单旋
			}else{
				T = DoubleRightLeftRotation(T)  // 右-左单旋
			}
		}
	}
	/* else val == T.Val 无需插入*/
	// 别忘了更新树高
	T.Height = Max(GetHeight(T.Left), GetHeight(T.Right)) + 1
	return  T
}
// GetHeight get height of Tree
func GetHeight(T *TreeNode) (height int) {
	if T == nil {
		return  0
	}
	left := GetHeight(T.Left)
	right := GetHeight(T.Right)

	return  Max(left, right)+1
}
// SingleLeftRotation 左单旋
func SingleLeftRotation(A *TreeNode) *TreeNode {
	/* 注意：A必须有一个左子结点B */
	/* 将A与B做左单旋，更新A与B的高度，返回新的根结点B */
	// B要成为新的root
	// 为了满足BST的性质, B的左子树要挂到A的右子树
	B := A.Left
	A.Left = B.Right
	B.Right = A
	A.Height = Max(GetHeight(A.Left), GetHeight(A.Right)) + 1
	B.Height = Max(GetHeight(B.Left), A.Height) + 1

	return B
}

// DoubleLeftRightRotation 左-右双旋
func DoubleLeftRightRotation(A *TreeNode) *TreeNode{
	/* 注意：A必须有一个左子结点B，且B必须有一个右子结点C */
	/* 将A、B与C做两次单旋，返回新的根结点C */

	/* 将B与C做右单旋，C被返回 */
	A.Left = SingleRightRotation(A.Left)
	/* 将A与C做左单旋，C被返回 */
	return SingleLeftRotation(A)
}
// SingleRightRotation 右单旋
func SingleRightRotation(A *TreeNode) *TreeNode {
	/* 注意：A必须有一个右子结点B */
	/* 将A与B做右单旋，更新A与B的高度，返回新的根结点B */
	// B要成为新的root
	// 为了满足BST的性质, B的左子树要挂到A的右子树
	B := A.Right
	A.Right = B.Left
	B.Left = A
	A.Height = Max(GetHeight(A.Left), GetHeight(A.Right)) + 1
	B.Height = Max(GetHeight(B.Right), A.Height) + 1

	return B
}

// DoubleRightLeftRotation 右-左双旋
func DoubleRightLeftRotation(A *TreeNode) *TreeNode {
	/* 注意：A必须有一个右子结点B，且B必须有一个左子结点C */
	/* 将A、B与C做两次单旋，返回新的根结点C */

	/* 将B与C做左单旋，C被返回 */
	A.Right = SingleLeftRotation(A.Right)
	/* 将A与C做右单旋，C被返回 */
	return SingleRightRotation(A)
}
