package leetcode641

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */

type Node struct {
	val  int
	next *Node
	last *Node
}

type MyCircularDeque struct {
	size      int
	capacity  int
	dummyHead *Node
	tail      *Node
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	dummyHead := Node{
		val:  -1,
		next: nil,
		last: nil,
	}
	return MyCircularDeque{
		capacity:  k,
		dummyHead: &dummyHead,
		tail:      &dummyHead,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	newNode := Node{
		val:  value,
		next: nil,
		last: nil,
	}
	dummyNext := this.dummyHead.next
	newNode.next = dummyNext
	newNode.last = this.dummyHead
	if dummyNext != nil {
		dummyNext.last = &newNode
	}else{
		this.tail = &newNode
	}
	this.dummyHead.next = &newNode
	this.size++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	newNode := Node{
		val:  value,
		next: nil,
		last: nil,
	}

	newNode.last = this.tail
	this.tail.next = &newNode
	this.tail = &newNode

	this.size++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.size == 0 {
		return false
	}
	deleteNode := this.dummyHead.next
	deleteNext := deleteNode.next
	deleteNode.next = nil
	deleteNode.last = nil
	this.dummyHead.next = deleteNext
	if deleteNext != nil {
		deleteNext.last = this.dummyHead
	}else{
		this.tail = this.dummyHead
	}
	this.size --

	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.size == 0 {
		return false
	}
	deleteNode := this.tail
	this.tail = deleteNode.last
	this.tail.next = nil
	deleteNode.last = nil
	this.size--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.size == 0 {
		return -1
	}
	if this.dummyHead.next == nil {
		return -1
	}
	return this.dummyHead.next.val
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.size == 0 {
		return -1
	}
	return this.tail.val
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.size == this.capacity
}
