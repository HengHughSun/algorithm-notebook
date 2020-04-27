package main

// https://pintia.cn/problem-sets/1211841066264109056/problems/1234055373040824320
import "fmt"


const Null = -1

type node struct {
	Element byte
	Left    int8
	Right   int8
}

var (
	N  int8
	T1 []node
	T2 []node
)

func main() {
	var (
		R1 int8
		R2 int8
	)
	R1 = -1
	R2 = -1
	// 建第一棵二叉树
	fmt.Scanf("%d\n", &N)
	if N != 0 {
		T1 = make([]node, N)
		R1 = BuildTree(T1)
	}
	// 建第二棵二叉树
	fmt.Scanf("%d\n", &N)
	if N != 0 {
		T2 = make([]node, N)
		R2 = BuildTree(T2)
	}
	// 判断是否同构
	if Isomorphic(R1, R2) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	//fmt.Printf("%#v\n", T1)
	//fmt.Printf("%#v\n", T2)
	//fmt.Println(R1)
	//fmt.Println(R2)
}

func BuildTree(T []node) (Root int8) {
	check := make([]int, N)
	var (
		Left    int8
		Right   int8
	)
	for i := int8(0); i < N; i++ {
		fmt.Scanf("%c %c %c\n", &T[i].Element, &Left, &Right)
		if Left != '-' {
			Left = Left - '0'
			check[Left] = 1
		} else {
			Left = Null
		}
		if Right != '-' {
			Right = Right - '0'
			check[Right] = 1
		} else {
			Right = Null
		}
		T[i].Left = Left
		T[i].Right = Right
	}
	//var i int8
	for Root = 0; Root < N; Root++ {
		if check[Root] != 1 {
			break
		}
	}
	//Root = i
	return Root
}

func Isomorphic(R1, R2 int8) bool {
	if R1 == Null  && R2 == Null { // both empty
		return true
	}
	//if R1 == Null  || R2 == Null {// one of them is empty
	//	return false
	//}
	if (R1 == Null && R2 != Null)  || (R1 != Null && R2 == Null) {// one of them is empty
		return false
	}
	if T1[R1].Element != T2[R2].Element { // root is different
		return false
	}

	if T1[R1].Left == Null && T2[R2].Left == Null { // both have not left subtree
		return Isomorphic(T1[R1].Right, T2[R2].Right)
	}
	if T1[R1].Left != Null && T2[R2].Left != Null && (T1[T1[R1].Left].Element == T2[T2[R2].Left].Element){
		// no need to swap the left and the right
		return ( Isomorphic(T1[R1].Left, T2[R2].Left) && Isomorphic(T1[R1].Right, T2[R2].Right))
	}else{
		// need to swap the left and the right
		return ( Isomorphic(T1[R1].Left, T2[R2].Right) && Isomorphic(T1[R1].Left, T2[R2].Right))
	}
	return true

}
