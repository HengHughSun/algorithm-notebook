package main

import "fmt"

// 单链表实现:

type MyLinkedList struct {
	val int
	next *MyLinkedList
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	tmp := this.next
	for i :=0; i < index && tmp != nil ; i ++ {
		tmp = tmp.next
	}
	if tmp != nil {
		return tmp.val
	}
	return -1
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	tmp := &MyLinkedList{val: val,}
	tmp.next = this.next
	this.next = tmp
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	tmp := this.next
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = &MyLinkedList{val: val}

}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	tmp := this

	for i :=0; i < index && tmp.next != nil ; i ++ {
		tmp = tmp.next
	}
	if tmp != nil {
		xx := &MyLinkedList{val: val}
		xx.next = tmp.next
		tmp.next = xx
	}
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	// this.check()
	if index == 0 {
		this.next = this.next.next
	}else{

		tmp := this

		for i :=0; i < index && tmp.next != nil ; i ++ {
			tmp = tmp.next
		}
		if tmp.next != nil {
			tmp.next = tmp.next.next
		}
	}

}

func (this *MyLinkedList) check(){
	l := this
	for l != nil {
		fmt.Print(l.val,"->")
		l = l.next
	}
	fmt.Println("all")
}
//["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]
//[[],[1],[3],[1,2],[1],[1],[1]]
//["MyLinkedList","addAtHead","addAtHead","addAtHead","addAtIndex","deleteAtIndex","addAtHead","addAtTail","get","addAtHead","addAtIndex","addAtHead"]
//[[],[7],[2],[1],[3,0],[2],[6],[4],[4],[4],[5,0],[6]]
func main() {
	 obj := Constructor();
	 obj.AddAtHead(7);
	obj.AddAtHead(2);
	obj.AddAtHead(1);
	 obj.AddAtIndex(3,0);
	obj.DeleteAtIndex(2);
	obj.AddAtHead(6);
	obj.AddAtTail(4);
	param_1 := obj.Get(4);
	fmt.Println(param_1)
	obj.AddAtHead(4);
	obj.AddAtIndex(5,0);
	obj.AddAtHead(6);

}
//["MyLinkedList","addAtHead","addAtHead","addAtHead","addAtIndex","deleteAtIndex","addAtHead","addAtTail","get","addAtHead","addAtIndex","addAtHead"]
//[[],[7],[2],[1],[3,0],[2],[6],[4],[4],[4],[5,0],[6]]

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

