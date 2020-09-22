package leetcode

import "math"

type MinStack struct {
	stack     []int
	miniStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:     []int{},
		miniStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	currentMin := this.miniStack[len(this.miniStack)-1]
	if x >= currentMin {
		this.miniStack = append(this.miniStack, currentMin)
	} else {
		this.miniStack = append(this.miniStack, x)
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.miniStack = this.miniStack[:len(this.miniStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.miniStack[len(this.miniStack)-1]
}
