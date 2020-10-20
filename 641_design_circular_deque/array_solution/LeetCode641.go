package leetcode

type MyCircularDeque struct {
	data  []int
	front int
	tail  int
}

func Constructor(k int) MyCircularDeque {

	return MyCircularDeque{
		data:  make([]int, k+1),
		front: 0,
		tail:  0,
	}
}

/** Checks whether the circular deque is empty or not. */
func (cd *MyCircularDeque) IsEmpty() bool {
	return cd.front == cd.tail
}

/** Checks whether the circular deque is full or not. */
func (cd *MyCircularDeque) IsFull() bool {
	return (cd.front-cd.tail+len(cd.data))%len(cd.data) == len(cd.data)-1
}

func (cd *MyCircularDeque) InsertFront(value int) bool {
	if cd.IsFull() {
		return false
	}

	cd.data[cd.front] = value
	cd.front++
	cd.front %= len(cd.data)
	return true
}

func (cd *MyCircularDeque) InsertLast(value int) bool {
	if cd.IsFull() {
		return false
	}

	cd.tail--
	cd.tail = (cd.tail + len(cd.data)) % len(cd.data)
	cd.data[cd.tail] = value
	return true
}

func (cd *MyCircularDeque) DeleteFront() bool {
	if cd.IsEmpty() {
		return false
	}
	cd.front--
	cd.front = (cd.front + len(cd.data)) % len(cd.data)
	return true
}

func (cd *MyCircularDeque) DeleteLast() bool {
	if cd.IsEmpty() {
		return false
	}
	cd.tail++
	cd.tail %= len(cd.data)
	return true

}

/** Get the front item from the deque. */
func (cd *MyCircularDeque) GetFront() int {
	if cd.IsEmpty() {
		return -1
	}
	return cd.data[(cd.front-1+len(cd.data))%len(cd.data)]
}

/** Get the last item from the deque. */
func (cd *MyCircularDeque) GetRear() int {
	if cd.IsEmpty() {
		return -1
	}
	return cd.data[cd.tail]
}
