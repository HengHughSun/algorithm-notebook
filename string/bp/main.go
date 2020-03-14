package main

import "fmt"

//  brute force 暴力算法
func bf(S string,T string) (pos int) {
	j := 0
	for j < len(S) {
	i := 0
	k := j
	for i < len(T){
		if T[i] == S[k] {
			i ++ 
			k ++
		}else {
			break
		}
	}
	if len(T) == i {
			return j
	}
	j ++ 
	}
	return -1
}

func main() {
	main := "bbcbbs.com" 
	son := "co"
	pos := bf(main,son )

	if  0 <= pos {
		fmt.Println(pos)
	}else{
		fmt.Println("not found")
	}
}