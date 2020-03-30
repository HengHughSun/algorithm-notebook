package main

import "fmt"

type PolyNode struct {
	coef int  // 系数
	expon int // 指数
	link *PolyNode
}

func main() {
	// 读入多项式1
	P1 := ReadPoly()
	// 读入多项式2
	P2 := ReadPoly()
	// 乘法运算并输出
	PP  := Mult(P1, P2)
	printPoly(PP)
	// 加法运算并输出
	PA  := Add(P1, P2)
	printPoly(PA)
}

/*因此我们需要设计几个函数
	读入多项式的函数
	运算乘法的函数
	运算加法的函数
	输出多项式的函数
 */

func ReadPoly() (*PolyNode) {
	var l int
	fmt.Scan(&l)
	var (
		c int
		e int
	)
	head := &PolyNode{}
	Rear := head
	for l > 0  {
		fmt.Scanf("%d%d", &c, &e)
		Rear = attach(c,e, Rear)
		l--
	}
	return head.link
}

func attach(c,e int, rear *PolyNode) *PolyNode{
	p := &PolyNode{}
	(*p).coef = c
	(*p).expon = e
	rear.link = p
	rear = p
	return  rear
}



func Add(p1,p2 *PolyNode) *PolyNode {
	t1, t2 := p1, p2
	P := &PolyNode{}
	rear := P
	for t2 != nil && t1 != nil {
		if t1.expon == t2.expon {
			c := t1.coef + t2.coef
			if c != 0 {
				tmp := &PolyNode{}
				tmp.coef = c
				tmp.expon = t1.expon
				rear.link = tmp
				rear = tmp
			}
			t2 = t2.link
			t1 = t1.link

		}else if t1.expon > t2.expon {
			rear.link = t1
			rear = t1
			t1 = t1.link
		}else {
			rear.link = t2
			rear = t2
			t2 = t2.link
		}
	}

	for t2 != nil {
		rear.link = t2
		rear = t2
		t2 = t2.link
	}

	for t1 != nil {
		rear.link = t1
		rear = t1
		t1 = t1.link
	}
	return  P.link
}
// Mult product
func Mult(p1,p2 *PolyNode) *PolyNode {
	/*
	系数相乘, 指数相加
	1. 将乘法运算转化为加法运算
	将p1当前项(ci,ei)乘p2多项式,再加到结果多项式里
	t1 = P1 t2 = P2
	P = &PolyNode{}
	for t2 != nil {
		rear = attach(t2.coef * t1.coef, t2.expon + t1.expon, rear)
		t2 = t2.link
	}
	2. 逐项插入
	将P1当前项(c1,e1)乘p2当前项(c2,e2),并插入到结果多项式中,
	关键是要找到插入位置
	*/
	if p1 == nil || p2 == nil {
		return nil
	}
	p := &PolyNode{}
	rear := p
	t1 := p1
	t2 := p2
	//1. 构建一个初始结果多项式
	for t2 != nil {
		rear = attach(t2.coef * t1.coef, t2.expon + t1.expon, rear)
		t2 = t2.link
	}

	//2. 将另外一个多项式插入
	t1 = t1.link
	for t1 != nil {
		t2 = p2
		rear = p
		for t2 != nil {
			c := t2.coef * t1.coef
			e := t2.expon + t1.expon
			// 判断插入的位置
			// 1. 遍历找到合适的位置
			for rear.link != nil && rear.link.expon > e {
				rear = rear.link
			}
			// 2 等于 系数相加  不等于构建新的结点并插入
			if  rear.link != nil && rear.link.expon == e {
				rear.link.coef += c
				if rear.link.coef == 0 {
					rear.link = rear.link.link
				}
			}else{
				tmp := &PolyNode{coef: c, expon: e,}
				tmp.link = rear.link
				rear.link = tmp
				// 为什么要加这句
				rear = rear.link
			}
			t2 = t2.link
		}
		t1 = t1.link
	}
	return  p.link

}


func printPoly(PA *PolyNode) {
	if PA == nil {
		fmt.Print("0 0\n")
		return
	}
	fmt.Print(PA.coef," ",PA.expon)
	PA = PA.link
	for PA != nil {
		fmt.Print(" ",PA.coef," ",PA.expon)
		PA = PA.link
	}
	fmt.Println()
}

/*测试用例
3 1 3 -1 2 1 1
3 3 3 1 2 -1 1
4 3 4 -5 2  6 1  -2 0
3 5 20  -7 4  3 1
 */
