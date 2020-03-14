package main

import "fmt"

// 魔术师发牌问题 单循环链表
// 要先创建一个长度为13的初始链表，然后按照特定顺序填值

type node struct {
	data int
	next *node
}

// Newnode 初始化结点
func Newnode(data int) *node {
	return &node{
		data: data,
	}
}

// 链表初始化
func createLinkList() *node {
	r := head
	for i := 1; i < cardNumber; i++ {
		s := Newnode(0)
		if head == nil { // 这里其实是无用的 go中head被初始化为了 各类型的 零型 这里head.data 已经被赋值为0
			head = s
		} else {
			r.next = s
		}
		r = s
	}
	// 最后尾执行头
	r.next = head

	return head
}

// 发牌顺序计算
func magician() {
	countNumber := 2
	p := head
	p.data = 1
	for {
		for j := 0; j < countNumber; j++ {
			p = p.next
			if p.data != 0 { // 该位置有牌的话，则下一个
				p = p.next
				j--
			}
		}
		if p.data == 0 {
			p.data = countNumber //
			countNumber++
			if countNumber == 14 {
				break
			}
		}
		// 纸牌 所以最大为13
	}
}

var cardNumber int
var head *node

func main() {
	head = &node{}

	fmt.Print("请输入牌个数: ")
	fmt.Scan(&cardNumber)
	createLinkList()
	magician()
	for i := 0; i < cardNumber; i++ {
		fmt.Printf("黑桃%d ", head.data)
		head = head.next
	}
}
