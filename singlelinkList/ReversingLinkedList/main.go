package main

import "fmt"

// https://pintia.cn/problem-sets/1211841066264109056/problems/1231458941036285954
const  MaxNum  = 1e5+10
type Node struct {
	id int
	data int
	next int
	//next *Node
}
j'j'j
jjjfunc main() {
	var (
		head int
		// N 长度
		N int
		// K 前K位翻转
		K int
	)
	fmt.Scanf("%d %d %d",&head, &N, &K)
	// 接受输入
	linklist := readInput(N)
	// 调整链表, 因为会出现一个结点根本不在链表上,所以要进行筛选
	var ans  []*Node
	for head != -1 {
		ans=append( ans,linklist[head])
		head = linklist[head].next
	}
	// 按要求翻转链表
	n := len(ans)
	//sortNode(linklist, head, N)
	//reverse(ans, K)
	// 注意理解题意ervey K 就要翻转一次
	for i := 0; i < n-n%K; i+=K { // 每 K 个结点一个区间
		for j := 0; j < K/2; j++{ // 反转链表
			t := ans[i+j];
			ans[i+j] = ans[i+K-j-1];
			ans[i+K-j-1] = t;
		}
	}
	// print
	for i := 0 ; i < n-1 ; i++ {
		fmt.Printf("%05d %d %05d\n",ans[i].id, ans[i].data, ans[i+1].id)
	}
	fmt.Printf("%05d %d -1\n",ans[n-1].id, ans[n-1].data)
}

func readInput(N int) (list []*Node) {
	/*
	00100 6 4
	00000 4 99999
	00100 1 12309
	68237 6 -1
	33218 3 00000
	99999 5 68237
	12309 2 33218
	*/
	list = make([]*Node, MaxNum)
	for i := 0; i < N; i ++ {
		var (
			id int
			data int
			next int
		)
		fmt.Scanf("%d %d %d",&id, &data, &next)
		list[id] = &Node{id:id,data:data,next:next}
		//list[i] = &Node{id:id,data:data,next:next}
	}
	return list
// 	head := &link{}
// 	tmp := head
// 	for start != -1 {
// 		tmp.next = &link{
// 			Node: list[start],
// 		}
// 		start = list[start].next
// 		tmp = tmp.next
// 	}
// 	head.next = reverse(head, K)
// 	tmp = head.next
// 	for tmp.next != nil {
// 		fmt.Printf("%05d %d %05d\n",tmp.id, tmp.data, tmp.next.id)
// 		tmp = tmp.next
// 	}
// 	fmt.Printf("%05d %d -1\n",tmp.id, tmp.data)
// 
}



// func reverse( head *link, K int) *link {
// 	cnt := 1
// 	new := head.next
// 	old := new.next

// 	for {
// 		for cnt < K {
// 			tmp := old.next
// 			old.next = new

// 			new = old
// 			old = tmp
// 			cnt++
// 		}
		
// 	}
// 	head.next.next = old
// 	return new
// }



//func sortNode(list []*Node, head, N int) {
//	for i := 0 ; i < N; i ++ {
//		for j := 0 ; j < N; j ++ {
//		if head == list[j].id {
//			head = list[j].next
//			list[i], list[j] = list[j], list[i]
//			break
//		}
//		}
//	}
//}

//#include<stdio.h>
//#include<iostream>
//#define MaxSize 100005
//using namespace std;
//int main(){
//int Data[MaxSize];
//int Next[MaxSize];
//int list[MaxSize];
//int FirstAdd,N,K;
//scanf("%d %d %d",&FirstAdd,&N,&K);
//for(int i=0;i<N;i++){
//  int tmpAdd,tmpData,tmpNext;
//  scanf("%d %d %d",&tmpAdd,&tmpData,&tmpNext);
//  Data[tmpAdd] = tmpData;
//  Next[tmpAdd] = tmpNext;
//}
//int sum=0;   // 累计有效结点数
//while(FirstAdd!=-1){   // 当尾结点为 -1 时结束
//  list[sum++] = FirstAdd;   // 记录所有Address
//  FirstAdd = Next[FirstAdd];  // 找下一个结点
//}
//for(int i=0;i<sum-sum%K;i+=K){  // 每 K 个结点一个区间
//  for(int j=0;j<K/2;j++){  // 反转链表
//      int t = list[i+j];
//      list[i+j] = list[i+K-j-1];
//      list[i+K-j-1] = t;
//  }
//}
//for(int i=0;i<sum-1;i++)
//  printf("%05d %d %05d\n",list[i],Data[list[i]],list[i+1]);
//printf("%05d %d -1\n",list[sum-1],Data[list[sum-1]]);
//return 0;
//}

//type Node struct {
//	address int // 结点自己的位置
//	data int
//	next int
//}
//type  Linknode struct {
//	node Node
//	nextnode *Linknode
//}
//func main() {
//	var head int
//	var (
//		N int // 长度
//		K int // 翻转的长度
//	)
//	fmt.Scanf("%d %d %d",&head, &N, &K)
//	// 接受输入
//	linklist := readInput(N)
//	// 调整链表
//	sortlist := sort(linklist, head, N)
//	// 按要求翻转链表
//
//
//
//	sortlist = sortlist.nextnode
//	for i := 0; i < N-1; i ++{
//		fmt.Printf("%05d %d %05d\n",sortlist.node.address, sortlist.node.data, sortlist.node.next)
//
//		sortlist = sortlist.nextnode
//	}
//	fmt.Printf("%05d %d -1\n",sortlist.node.address, sortlist.node.data)
//	// 输出链表
//}
//
//func readInput(l int) *Linknode {
//	head := &Linknode{}
//	rear := head
//	for i := 0; i < l; i++ {
//		tmp := &Linknode{}
//		fmt.Scanf("%d %d %d",&tmp.node.address, &tmp.node.data, &tmp.node.next)
//		rear.nextnode = tmp
//		rear = tmp
//	}
//	return head.nextnode
//}
//
//func sort(list *Linknode, head int, N int) *Linknode {
//	p := &Linknode{}
//	newrear := p
//	j := N
//	for j > 0{
//		rear := list
//		for i := 0;  i < N; i++ {
//			if rear != nil {
//				if rear.node.address == head {
//					// 构建新链表
//					tmp := &Linknode{}
//					tmp.node = rear.node
//					newrear.nextnode = tmp
//					newrear = newrear.nextnode
//					// 更新head
//					head = rear.node.next
//					j--
//				}
//				rear = rear.nextnode
//			}
//		}
//	}
//	return p
//}
//
//func reverse(list *Linknode,N, K int) *Linknode {
//	rear := list.nextnode
//	head := rear
//
//	for i := 0; i < K ; i++ {
//		rear = rear.nextnode
//	}
//	head.node.next = rear.node.next
//	/*
//	00000 4 33218 00100 1 12309
//	33218 3 12309 12309 2 33218
//	12309 2 00100 33218 3 00000
//	00100 1 99999 00000 4 99999
//	99999 5 68237
//	68237 6 -1
//	 */
//
//}