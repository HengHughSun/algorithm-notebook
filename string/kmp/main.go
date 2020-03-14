package main

import "fmt"

func getNext(T string) (next []int) {
	pLen := len(T)
	next = make([]int, pLen)
	i := 1
	j := 0
	// next[0] = len(T)
	next[0] = -1
	for i < pLen-1 {
		if 0 == j || T[i] == T[j] {
			i++
			j++
			// 优化算法
			if T[i] != T[j] { // T 到3就结束了
				next[i] = j
			} else {
				next[i] = next[j]
			}
			/*
			 0   1  2  3
			 -1  0  1  2
			*/
		} else {
			j = next[j]
		}
	}
	return

}

// IndexKMP 返回字串T在主串S第pos个位置之后的值
func IndexKMP(S string, T string) (pos int) {
	i := 0
	j := 0
	k := getNext(T)
	for i < len(S) && j < len(T) {
		if -1 == j || S[i] == T[j] {
			i++
			j++
		} else {
			j = k[j]
			// if -1 == j {
			// 	j++
			// 	i++
			// }
		}
	}
	if len(T) == j {
		return i - j // j此时就==len(T)
	}
	return -1
}

func main() {
	main := "bbcbbs.com"
	son := "bbbc"

	pos := IndexKMP(main, son)
	if 0 <= pos {
		fmt.Println(pos)
	} else {
		fmt.Println("not found")
	}
}
