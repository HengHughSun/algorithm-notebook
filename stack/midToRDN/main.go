package main

import (
	"fmt"
)

// binary to decimal 二进制转化为十进制

const (
	// StackInitSize Initial size
	StackInitSize = 20
)

type eleMent byte

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
	}
	e = s.Data[s.Top]
	s.Top--
	return e
}

func (s sqStack) len() int {
	if s.Top == -1 {
		return 0
	}
	return s.Top + 1
}

func priority(s *sqStack, e rune) {
	// fmt.Print(eleMent(e), " ")
	for !s.IsEmpty() && s.Data[s.Top] >= eleMent(e) {
		fmt.Println(string(s.pop()), " ")
	}
	s.push(eleMent(e))
}

var name []rune

func main() {
	name = []rune{'1', '+', '(', '2', '-', '3', ')', '*', '4', '+', '9', '/', '5', '#'}
	// name = []string{"1", "+", "(", "2", "-", "3", ")", "*", "4", "+", "10", "/", "5", "#"}

	stack := initialStack()
	fmt.Print("请输入后缀表达式,数字之间用空格隔开,输入#符号表示结束! \n")
	// fmt.Scan(&name)
	// inputReader := bufio.NewReader(os.Stdin)
	// name, err := inputReader.ReadBytes('#')
	// if err != nil {
	// 	fmt.Printf("输入出错, error: %s\n", err)
	// }
	for _, e := range name {

		switch e {
		case '#':
			break
		case ' ':
			continue
		case '+', '-':
			priority(stack, e)
		case '*', '/':
			for !stack.IsEmpty() && stack.Data[stack.Top] <= eleMent(e) {
				fmt.Print("间接输出", string(stack.pop()), " ")
			}
		case '(':
			stack.push(eleMent(e))
		case ')':
			for !stack.IsEmpty() && stack.Data[stack.Top] != eleMent('(') {
				fmt.Print("间接输出", string(stack.pop()), " ")
			}
			stack.pop()
		default:
			fmt.Println(string(e), " ")
		}
	}
	len := stack.len()
	fmt.Printf("栈的当前容量是: %d", len)
	for !stack.IsEmpty() {
		fmt.Print("剩余输出", string(stack.pop()), " ")
	}
}
