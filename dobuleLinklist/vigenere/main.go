package main

import (
	"fmt"
	"math/rand"
)

/*
Vigenere 维吉尼亚加密:
- 当输入明文,自动生成随机密钥匹配明文中每个字母并移位加密
明文     I L  O  V  E
随机密匙 3 15 23 2  52
密文	 L A  L  X  E
*/
type node struct {
	data  byte
	prior *node
	next  *node
}

// Newnode _
func Newnode(data byte) *node {
	return &node{
		data: data,
	}
}

func createLinkList() {
	p := head
	for i := byte(0); i < 26; i++ {
		q := Newnode('A' + i)
		q.next = p.next
		q.prior = p
		p.next = q
		p = q
	}
	p.next = head.next
	head.next.prior = p
}

func cryptoByte(b byte, i int) byte {
	p := head.next
	tmp := key[i] % 26

	for {
		p = p.next
		if p.data == b {
			for tmp > 0 {
				p = p.next
				tmp--
			}

			break
		}
	}
	return p.data
}
func uncryptoByte(b byte, i int) byte {
	p := head.next
	tmp := key[i] % 26
	for {
		p = p.prior
		if p.data == b {
			for tmp > 0 {
				p = p.prior
				tmp--
			}

			break
		}

	}
	return p.data
}

var head *node

// S __
var S string
var key []int

func main() {
	r := rand.New(rand.NewSource(99))
	head = &node{}
	fmt.Scan(&S)
	l := len(S)
	s := []byte(S)
	key = make([]int, l)
	for i := 0; i < l; i++ {
		key[i] = r.Intn(99)
	}
	createLinkList()
	res := make([]byte, l)
	for i := 0; i < l; i++ {
		res[i] = cryptoByte(s[i], i)
	}
	fmt.Println(key)
	fmt.Println(string(res))
	unres := make([]byte, l)
	for i := 0; i < l; i++ {
		unres[i] = uncryptoByte(res[i], i)
	}
	fmt.Println(string(unres))
}
