package leetcode641

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestStruct(t *testing.T){
	deque :=Constructor(3)
	
	deque.InsertLast(1)
	deque.InsertLast(2)
	deque.InsertFront(3)
	deque.InsertFront(4)
	rear := deque.GetRear()

	
	assert.Equal(t,2,rear)

}