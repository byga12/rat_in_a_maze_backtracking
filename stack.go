package main

type Stack[T any] struct {
	top *Node[T]
	length uint
}

func NewStack[T any]() *Stack[T]{
	return &Stack[T]{nil,0} 
}

func (stack *Stack[T]) push(value T) {
	node := &Node[T]{value: value}
	if stack.length == 0 && stack.top == nil {
		stack.top = node
	} else {
		node.next = stack.top
		stack.top = node
	}
	stack.length++
}

func (stack *Stack[T]) pop(){
	if stack.length == 0 {return}
	if stack.length == 1 {
		stack.top = nil
	} else {
		// De limpiar el nodo anterior se encarga el garbage collector ?
		stack.top = stack.top.next
	}
	stack.length --
}

func(stack *Stack[T]) isEmpty() bool{
	if stack.length == 0 {
		return true
	} else {
		return false
	}
}