package main

import "fmt"

// half-interval search  二分查找
//
func half(a []int, target int) {
	left, right := 0, len(a)-1
	mid := (left + right) / 2

	if a[mid] == target {
		fmt.Println(mid)
	} else if target < a[mid] {

		half(a[:mid], target)
	} else {
		half(a[mid+1:], target)
	}

}

func main() {
	var a []int
	a = []int{1, 4, 5, 8, 9}
	half(a, 8)

}
