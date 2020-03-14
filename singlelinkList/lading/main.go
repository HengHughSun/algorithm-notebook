package main

import "fmt"

// 拉丁方阵问题  n x n 方阵问题
/*
形如:
	1  2  3
	2  3  1
	3  1  2
要求N由用户输入

*/
type node struct {
	data int
	next *node
}

// Newnode __
func Newnode(data int) *node {
	return &node{
		data: data,
	}
}

func createLinkList() {
	r := head
	for i := 2; i <= N; i++ {
		s := Newnode(i)
		r.next = s
		r = s
	}

	r.next = head
}

var head *node

// N _
var N int

func main() {
	fmt.Print("请输入N的值: ")
	fmt.Scan(&N)
	head = &node{data: 1}
	createLinkList()

	k := head
	for i := 0; i < N; i++ {
		fmt.Printf("第%d行是:\n", i)
		x := k
		for {
			fmt.Print(k.data, " ")
			k = k.next
			if k.data == x.data {
				break
			}
		}
		k = k.next
		fmt.Print("\n")
	}

}
