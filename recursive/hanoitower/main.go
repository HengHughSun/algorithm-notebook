package main

import "fmt"

// 汉诺塔
// 要求输出移动步骤
var s int

func main() {
	n := 3 //盘子数
	Hanoi(n, "A", "B", "C")
	fmt.Println("总共", s)
}

//Hanoi 将src上的盘子借助mid移动到dest
func Hanoi(n int, src, mid, dest string) {
	// 将src的n个盘子以 mid为中转,移动到dest
	if n == 1 {
		fmt.Println("n=", n, " ", src, "-->", dest)
		s++
		return
	}
	Hanoi(n-1, src, dest, mid)                  // 先将n-1个盘子从src借助dest移动到 mid
	fmt.Println("n=", n, " ", src, "-->", dest) //将第n个盘子从src移动到dest上
	Hanoi(n-1, mid, src, dest)                  // 再将n-1个盘子从mid借助src移动到 dest

}
