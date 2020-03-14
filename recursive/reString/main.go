package main

import "fmt"
// 翻转字符串
func print(i int) {
		if a[i] != '#' {
			print(i+1)
		}
		if a[i] != '#' {
			fmt.Print(string(a[i]))
		}
}
var a []byte
func main() {
	fmt.Scan(&a)
	i := 0
		print(i)
}