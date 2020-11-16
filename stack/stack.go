package stack

import ()

// Stack: Container that has LIFO order.

type stack struct {
	top 	*node
	length 	int
}

type node struct {
	data 	interface{}
	next 	*node
}

// NewStack creates a new stack container
func NewStack() *stack {
	return &stack{nil, 0}
}

// IsEmpty checks if the stack is empty or not
func (st *stack) IsEmpty() bool {
	return (st.top == nil)
}

// Size return the number of elements in the stack
func (st *stack) Size() int {
	return st.length
}

// Top return the value of the top element in the stack
func (st *stack) Top() interface{} {
	if st.top == nil {
		return nil
	}

	return st.top.data
}

// Push insert the given element into the stack
func (st *stack) Push(d interface{}) {
	newNode := &node{d, st.top}
	st.top = newNode
	st.length++
}

// Pop removes the top element from the stack
func (st *stack) Pop() {
	if st.top == nil {
		return
	}

	st.top = st.top.next
	st.length--
}