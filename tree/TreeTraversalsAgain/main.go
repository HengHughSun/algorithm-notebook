package main

import (
	"fmt"
)

// 不建树的操作
type stack struct {
	data []int
}

func Newstack() *stack {
	return &stack{
		data: make([]int, 0),
	}
}

func (s *stack)Push(val int) {
	s.data = append(s.data, val)
}

func (s *stack)Pop() (val int) {
	val = s.data[len(s.data) - 1]
	s.data = s.data[:len(s.data) - 1]
	return val
}

func main() {
	var pre []int
	var in  []int

	var l  int
	fmt.Scanf("%d\n", &l)

	pre = make([]int,0, l)
	in = make([]int, 0, l)
	// 读取输入 建立 先序遍历pre   中序遍历in

	stack := Newstack()
	for i := 0; i < 2*l ; i++ {
		var state string
		fmt.Scanf("%s", &state)
		switch state {
		case "Push":
			var value int
			fmt.Scanf("%d", &value)
			pre = append(pre, value)
			stack.Push(value)
		case "Pop":
			in = append(in, stack.Pop())
		}
	}

	var Post []int
	Post = make([]int, l)
	slove(pre, in,Post, l)

	var first = true 
    for i := 0; i < l; i ++ {
        if first{
            fmt.Print(Post[i])
            first = false
	    }else{

		fmt.Print(" ",Post[i])
	    }
    }
}

// slove Prel 前序 In 中序  n 数结点数
func slove(Prel []int, In []int, Post []int, l int){
	if l == 0 { return }
	if l == 1 {
		Post[len(Post)-1] = Prel[0]
		return
	}
	Post[len(Post)-1] = Prel[0]
	root := Prel[0]

	var i int
	for i = 0; In[i] != root; i ++ {}

	slove(Prel[1: 1 + i], In[:i],Post[:i], i)
	slove(Prel[i+1:], In[i + 1:],Post[i:l-1], l-i-1)

}