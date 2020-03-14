package main

import (
	"bufio"
	"fmt"
	"os"
)

// binary to decimal 二进制转化为十进制

const (
	// StackInitSize Initial size
	StackInitSize = 20
)

type eleMent float64

// 结点
type sqStack struct {
	Data    []eleMent // 	存储元素的数组
	Top     int       //  栈顶指针
	MaxSize int       // 堆栈最大容量
}

func initialStack() *sqStack {
	Data := make([]eleMent, StackInitSize)
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

func (s *sqStack) push(e eleMent) {
	if s.IsFull() {
		s.Data = append(s.Data, e)
		// 扩容
	} else {
		s.Top++
		s.Data[s.Top] = e
	}
}

func (s *sqStack) pop() (e eleMent) {
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

var name []eleMent

func main() {
	// name = []eleMent{'1', '2', '-', '4', '5', '+', '*', '#'}

	stack := initialStack()
	fmt.Print("请输入后缀表达式,数字之间用空格隔开,输入#符号表示结束! \n")
	// fmt.Scan(&name)
	inputReader := bufio.NewReader(os.Stdin)
	name, err := inputReader.ReadBytes('#')
	if err != nil {
		fmt.Printf("输入出错, error: %s\n", err)
	}
	for _, e := range name {
		switch e {
		case '#':
			break
		case ' ':
			continue
		case '+':
			stack.push(stack.pop() + stack.pop())
		case '-':
			v := stack.pop()
			b := stack.pop()
			stack.push(b - v)
		case '*':
			stack.push(stack.pop() * stack.pop())
		case '/':
			v := stack.pop()
			b := stack.pop()
			if v != 0 {
				stack.push(b / v)
			} else {
				fmt.Errorf("出错:除数为%d", b)
			}
		default:
			stack.push(eleMent(e) - 48)
		}
	}
	len := stack.len()
	// stack8 := initialStack()
	fmt.Printf("栈的当前容量是: %d\n", len)
	fmt.Printf("值为%f\n", stack.pop())

	len = stack.len()
	fmt.Printf("栈的当前容量是: %d", len)
}
