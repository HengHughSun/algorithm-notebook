package main

/*
https://pintia.cn/problem-sets/1211841066264109056/problems/1234055373045018625
如果对二叉树的遍历非递归实现有很好理解的话，那么这里理解起来就不是很大的难度，在输入格式中Push的过程就是二叉树前序遍历的结果，
而Pop的过程及为二叉树的中序遍历的结果。如果理解了这两个操作，那么题目的意思就变成了已知二叉树的前序遍历和中序遍历求二叉树的后序遍历。
由于前序遍历首先访问的是根节点，中序遍历首先访问的是左子树，而后序遍历最后才访问根节点。因此，前序遍历中第一个元素为后序遍历的最后一
个元素，在中序遍历中根节点左边的为左子树的节点，右边为右子树的节点。所以需要递归的求解这个问题。
原文链接：https://blog.csdn.net/shinanhualiu/article/details/49279051
题目实质是通过先序遍历和中序遍历建树，再后序遍历树。
https://www.cnblogs.com/llhthinker/p/4748792.html
 */
import "fmt"

var N int

type Node struct{
	data int
	left *Node
	right *Node
}

type Stack struct {
	data []*Node
	top int
}
func (s *Stack)Push(node *Node) {
	s.data = append(s.data,node)
	s.top++
}
func (s *Stack)Pop() *Node {
	tmp := s.data[s.top]
	s.data = s.data[:s.top]
	s.top--
	return tmp
}

func (s *Stack)IsEmpty() bool {
	return s.top == -1
}

func (s *Stack)Top() *Node {
	return s.data[s.top]
}

func main() {
	fmt.Scanf("%d\n", &N)
	T := &Stack{
		top:  -1,
	}
	buildtree(T)
	postreversal(root)
}
var first = true
func postreversal(root *Node) {
	head := root
	if head == nil {
		return
	}
	postreversal(head.left)
	postreversal(head.right)
	if first{
		fmt.Print(head.data)
		first = false
	}else{

		fmt.Print(" ",head.data)
	}
}


func NewNode(data int) *Node{
	return &Node{data: data}
}
var root *Node
var father *Node
func buildtree(T *Stack){
	var (
		action string
		data int
		findroot = false
		poped = false
	)

	for i := 0; i < 2*N; i++ {
		fmt.Scanf("%s",&action)
		if action == "Push" {
			fmt.Scanf("%d",&data)
			newnode := NewNode(data)
			if !findroot {
				root = newnode
				T.Push(root)
				findroot = true
			}else {
				if !poped {
					father = T.Top()
				}
					if father.left == nil {
						father.left = newnode
					}else {
						father.right = newnode
					}
				T.Push(newnode)
			}
				poped = false
		}else {
				father = T.Pop()
			poped = true
		}
	}
}
