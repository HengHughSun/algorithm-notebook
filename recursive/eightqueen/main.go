package main

import "fmt"

// N 皇后问题
var N int
var s int
var queenPos []int

func main() {
	// fmt.Scan(&N)
	N = 8
	queenPos = make([]int, N)
	//for i := 0; i < len(queenPos); i++ {
	//	queenPos[i] = 1
	//}
	Nqueen(0)
	fmt.Println("多少种摆法", s/N)
}

// Nqueen 假设0~k行的皇后都已经拜访好了
func Nqueen(k int) { // 在0~k-1行皇后已经摆好的情况下，摆第K行及k行后的皇后
	var i int   // 在这里坑了许久
	if k == N { // N个皇后已经摆好了
		for i = 0; i < N; i++ {
			fmt.Println("位置", queenPos[i])
			s++
		}
		return

	}
	for i = 0; i < N; i++ { // 尝试当前第K个皇后的位置
		var j int
		for j = 0; j < k; j++ {
			// 和已经摆好的k个皇后的位置比较，看是否冲突
			if queenPos[j] == i || Abs(queenPos[j]-i) == Abs(k-j) { // Abs(queenPos[j]-i)  这里i写成了1 坑了很长时间
				break // 冲突，测试下一个位置
			}
		}
		if j == k { // 当前选的位置i 不冲突
			queenPos[k] = i // 将第k个皇后放在该位置
			// fmt.Println(queenPos)
			Nqueen(k + 1)
		}
		// fmt.Println(queenPos)
	}

}

// Abs _
func Abs(n int) int {
	if n < 0 {
		// fmt.Println("绝对值", n, -n)
		return -n
	}
	return n

}
