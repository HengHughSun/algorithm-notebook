package main

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
