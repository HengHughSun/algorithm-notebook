package main

import (
	"fmt"
)

// 线索存储标志位
// link(0)L 表示指向左右孩子的指针
// Thread(1): 表示指向前驱后继的线索
const (
	Link = iota
	Thread
)

// pre 始终指向刚刚访问过的结点 全局变量
var pre *BiTNode
// byte define tmp type

// BiTNode ..
type BiTNode struct {
	lchild *BiTNode
	ltag   int
	data   byte
	rchild *BiTNode
	rtag   int
}

// CreateBiTree 创建一颗二叉树，约定用户遵照前序遍历
// 前序构建树 是最方便的
func CreateBiTree(BiTree *BiTNode) {
	var c byte
	fmt.Scanf("%c", &c)
	if ' ' == c {
		BiTree = nil
	} else {
		BiTree.data = byte(c)
		BiTree.ltag = Link
		BiTree.rtag = Link
		BiTree.lchild = new(BiTNode)
		BiTree.rchild = new(BiTNode)
		//BiTree.lchild = &BiTNode{}
		//BiTree.rchild = &BiTNode{}

		CreateBiTree((*BiTree).lchild)
		CreateBiTree((*BiTree).rchild)
	}
}

// func (n *Node) Insert(value, data string) error {

// 	if n == nil {
// 		return errors.New("Cannot insert a value into a nil tree")
// 	}
// 	switch {
// 	case value == n.Value:
// 		return nil
// 	case value < n.Value:
// 		if n.Left == nil {
// 			n.Left = &Node{Value: value, Data: data}
// 			return nil
// 		}
// 		return n.Left.Insert(value, data)
// 	case value > n.Value:
// 		if n.Right == nil {
// 			n.Right = &Node{Value: value, Data: data}
// 			return nil
// 		}
// 		return n.Right.Insert(value, data)
// 	}
// 	return nil
// }

// visit 访问二叉树结点的具体操作，你想干嘛？
func visit(T byte) {
	fmt.Printf("%c", T)
}

// PreOrderTraversal  前序遍历二叉树
//func PreOrderTraversal(BinaryTree *BiTNode, level int) {
//	if BinaryTree != nil {
//		visit(BinaryTree.data, level)
//		PreOrderTraversal(BinaryTree.lchild, level+1)
//		PreOrderTraversal(BinaryTree.rchild, level+1)
//	}
//}

// InThreading 中序遍历进行线索化
func InThreading(T *BiTNode) {
	if T.lchild != nil || T.rchild != nil {
		InThreading(T.lchild) // 递归左孩子线索化,
		// 结点处理
		if T.lchild.lchild == nil { // 如果该结点没有左孩子，设置ltag为Thread,并把lchild指向刚刚访问的结点
			T.ltag = Thread
			T.lchild = pre
		}
		if pre.rchild.rchild == nil {
			pre.rtag = Thread
			pre.rchild = T
		}
		pre = T
		InThreading(T.rchild) // 递归右孩子线索化
	}
}


// InOrderThreading ...
func InOrderThreading(P *BiTNode, T *BiTNode) {
	P.ltag = Link
	P.rtag = Thread
	P.rchild = P

	if T == nil {
		P.lchild = P
	} else {
		P.lchild = T
		pre = P
		InThreading(T)
		pre.rchild = P
		pre.rtag = Thread
		P.rchild = pre
	}
}

//
// InOrderTraversal  中序遍历二叉树
func InOrderTraversal(head *BiTNode) {
	var tmp *BiTNode
	tmp = head.lchild  // 'A'
	for tmp != head {
		for tmp.ltag == Link {
			if tmp.lchild != nil {
				tmp = tmp.lchild
			}
		}
		visit(tmp.data)
		for tmp.rtag == Thread && tmp.rchild != head {
			tmp = tmp.rchild
			visit(tmp.data)
		}
		tmp = tmp.rchild

	}
}
func main() {
	// level := 0
	var (
		T *BiTNode
		PRE *BiTNode
	)
	T = new(BiTNode)
	PRE = new(BiTNode)
	CreateBiTree(T)
	InOrderThreading(PRE, T)
	// PreOrderTraversal(&T, level)
	fmt.Print("中序遍历输出结果为:\n")
	InOrderTraversal(PRE)

}
