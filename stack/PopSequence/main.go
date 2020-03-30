package main

import "fmt"

type stack struct {
	top int
	data []int
}

// https://pintia.cn/problem-sets/1211841066264109056/problems/1231458941036285955

func Newstack() *stack{
	return &stack{
		top: 0,
		data: []int{0,},
	}
}

func (s *stack)push(val int) {
	s.data = append(s.data, val)
	s.top++
}

func (s *stack)pop() (val int) {
	val = s.data[s.top]
	s.data = s.data[:s.top]
	s.top--
	return
}
var (
	M  int
	N  int
	K  int // 要检查的pop长度
)

func main() {
	fmt.Scanf("%d %d %d",&M, &N, &K)

	vec := make([]int,0)

	for i := 0; i < K; i++ {
		for j := 0; j < N; j ++ {
			var val int
			fmt.Scanf("%d",&val)
			vec = append(vec, val)
		}
		if check(vec) {
			fmt.Println("YES")
		}else{
			fmt.Println("NO")
		}
		vec = vec[:0]
	}

}


func check(ori []int) bool {
	sta := Newstack()
	i := 0
	num := 0
	for i < N {
		//sta.push(ori[i])
		for ori[i] > sta.data[sta.top] && sta.top < M  { // ori[i] 想要pop出, 必须满足小于他的都进过栈
			/*
			7 5 5
			4 3 2 1 5 6 7
			4 要被第一个pop出来, 3 2 1就必须在栈中
			 */
			num ++
			sta.push(num)
		}
		if ori[i] != sta.data[sta.top] {
			return  false
		}else{
			sta.pop()
		}
		i++
	}

	return true
}
