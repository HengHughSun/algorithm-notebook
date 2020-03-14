package main

import (
	"bufio"
	"fmt"
	"os"
)

// 凯撒密码 双向循环链表
/*
要求实现用户输入一个数使得26个字母的排列发生变化，同时支持负数，
例如用户输入3
其结果为: DEFGHIJKLMNOP
*/
var head *node

// byte unit8
type node struct {
	data  byte
	prior *node
	next  *node
}

// Newnode 构造结点
func Newnode(data byte) *node {
	return &node{
		data: data,
	}
}

// createLinkList  p 为当前尾结点的指针   q为要插入结点的指针
func createLinkList() {
	p := head
	for i := uint8(0); i < 26; i++ {
		q := Newnode('A' + i)
		q.next = p.next
		p.next = q
		q.prior = p
		p = q
	}
	p.next = head.next
	head.next.prior = p
}

func caeser(init int) {
	if init > 0 {
		for init > 0 {
			head = head.next
			init--
		}

	}
	if init < 0 {
		head = head.next
		for init <= 0 {
			head = head.prior
			init++
		}

	}
}

// N _
var N int

func main() {
	head = &node{}
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入偏移量: ")
	tmp1, err := fmt.Fscan(inputReader, &N)
	if err != nil {
		fmt.Print(tmp1)
		fmt.Print(err)
	}
	createLinkList()
	caeser(N)
	tmp := head.next
	for {
		fmt.Print(string(tmp.data))
		tmp = tmp.next
		if tmp.data == head.next.data {
			break
		}
	}

}
