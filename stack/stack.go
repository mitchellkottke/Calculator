// Package Stack is a standard stack

package stack

type (
	node struct {
		next  *node
		value any
	}

	Stack struct {
		head *node
	}
)

// Create an empty stack
func New() *Stack {
	return &Stack{head: nil}
}

// Push an item on the stack
func (stack *Stack) Push(item any) {
	stack.head = &node{next: stack.head, value: item}
}

// Pop an item off the stack
func (stack *Stack) Pop() any {
	if stack.head == nil {
		return nil
	}

	value := stack.head.value
	stack.head = stack.head.next
	return value
}

// Look at the top item of the stack
func (stack *Stack) Peek() any {
	if stack.head == nil {
		return nil
	}

	return stack.head.value
}
