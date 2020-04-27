package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d\n", &N)
	var S []int
	S = make([]int, N)
	// Initialization
	for i := 0; i < N; i++ {
		S[i] = -1
	}
	var In rune
	for In != 'S' {
		fmt.Scanf("%c", &In)
		switch In {
		case 'I':
			Input_connection(S)
		case 'C':
			Check_connection(S)
		case 'S':
			Check_network(S, N)
		}
	}
}

func Check_network(S []int, n int) {
	var counter = 0
	for i := 0; i < n; i++ {
		if S[i] < 0 {
			counter++
		}
	}
	if counter == 1 {
		fmt.Println("The network is connected.")
	} else {
		fmt.Printf("There are %d components.\n", counter)
	}
}

func Check_connection(S []int) {
	var r1, r2 int

	fmt.Scanf("%d %d\n", &r1, &r2)
	Root1 := Find(S, r1-1)
	Root2 := Find(S, r2-1)
	if Root2 == Root1 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func Input_connection(S []int) {
	var r1, r2 int

	fmt.Scanf("%d %d\n", &r1, &r2)
	Root1 := Find(S, r1-1)
	Root2 := Find(S, r2-1)

	if Root1 != Root2 {
		Union(S, Root1, Root2)
	}
}

func Union(S []int, root1 int, root2 int) {
	// 树2 高于 树1 小树贴到大树上
	// 比树高
	//if S[root2] < S[root1] {
	//	S[root1] = root2
	//	//S[root2] -= 1 // 同时树高是不改变的 所以这里不需要
	//}else{
	//	if S[root2] ==  S[root1] { //树高++
	//		S[root1]--
	//	}
	//	S[root2] = root1
	//}

	// 比规模

	if S[root2] < S[root1] {
		S[root2] += S[root1]
		S[root1] = root2
	} else {
		S[root1] += S[root2]
		S[root2] = root1
	}

}

func Find(S []int, i int) int {
	if S[i] < 0 {
		return i
	} else {
		S[i] = Find(S, S[i])
		return S[i]
	}
}
