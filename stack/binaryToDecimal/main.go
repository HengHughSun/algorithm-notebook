package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// binary to decimal 二进制转化为十进制

const (
	// StackInitSize Initial size
	StackInitSize = 20
	// StackIncrement increase
	StackIncrement = 10
)

// 结点
type sqStack struct {
	Data    []byte // 	存储元素的数组
	Top     int    //  栈顶指针
	MaxSize int    // 堆栈最大容量
}

func initialStack() *sqStack {
	Data := make([]byte, StackInitSize)
	return &sqStack{
		Data:    Data,
		Top:     -1,
		MaxSize: StackInitSize,
	}
}

func (s *sqStack) IsFull() bool {
	return s.Top == s.MaxSize-1
}

func (s *sqStack) IsEmpty() bool {
	return s.Top == -1
}

func (s *sqStack) push(e byte) {
	if s.IsFull() {
		s.Data = append(s.Data, e)
		// 扩容
	} else {
		s.Top++
		s.Data[s.Top] = e
	}
}

func (s *sqStack) pop() (e byte) {
	if s.IsEmpty() {
		fmt.Println("栈已空")
		return '0'
	} else {
		e = s.Data[s.Top]
	}
	s.Top--
	return e
}

func (s sqStack) len() int {
	if s.Top == -1 {
		return 0
	}
	return s.Top + 1
}

var name []byte

func main() {
	// name = []byte{'1', '1', '0', '1', '0'}

	// var i, sum = 0, 0

	stack := initialStack()
	fmt.Print("请输入二进制数,输入#符号表示结束! \n")
	// fmt.Scan(&name)
	inputReader := bufio.NewReader(os.Stdin)
	name, err := inputReader.ReadBytes('#')
	if err != nil {
		fmt.Printf("输入出错, error: %s\n", err)
	}
	for _, e := range name {
		if e == '#' {
			break
		}
		stack.push(e)
	}

	len := stack.len()
	fmt.Printf("栈的当前容量是: %d\n", len)
	sum := 0
	for i := 0; i < len; i++ {
		k := stack.pop()
		sum += (int(k) - 48) * int(math.Pow(2, float64(i)))

	}
	fmt.Println(sum)
	if stack.IsEmpty() {
		fmt.Println("空")
	}

	len = stack.len()
	fmt.Printf("栈的当前容量是: %d", len)
}
