package Containers
/*package main

import (
    "fmt"
)*/

type Element struct {
	value interface{}
	next  *Element
	prev *Element
}

type Container struct {
	top *Element
	bottom *Element
	size int
	level int
}

func (q *Container) GetLevel() (int) {
    return q.level
}

func (q *Container) AddLevel(){
    q.level++
}


// Remove the top element from the stack in first in first out order.
// If the stack is empty, return nil
func (q *Container) PopTop() (value interface{}) {
	if q.size > 0 {
		value, q.top = q.top.value, q.top.next
		q.size--
		return
	}
	return nil
}

// Pop removes and returns a Element from the Container in first to last order.
// If the stack is empty, return nil
func (q *Container) PopBottom() (value interface{}) {
	if q.size > 0 {
		value, q.bottom = q.bottom.value, q.bottom.prev
		q.size--
		return
	}
	return nil
}

func (q *Container) PeekTop() (value interface{}) {
	if q.size > 0 {
		return q.top.value
	}
	return nil
}
func (q *Container) PeekBottom() (value interface{}) {
	if q.size > 0 {
		return q.bottom.value
	}
	return nil
}


func (q *Container) Len() int{
    return q.size
}


// Push adds a Element to the Container.
func (q *Container) Push(value interface{}) {
    if q.top == nil {
        q.bottom = nil
	    q.top = &Element{value, q.bottom, nil}
	    q.size++
    } else if q.bottom == nil {
        q.top.next = &Element{value, nil, q.top}
        q.bottom = q.top.next
        q.size++
    } else {
        q.bottom.next = &Element{value, nil, q.bottom}
        q.bottom = q.bottom.next
        q.size++
    }
}


/*
func main() {
    stack := new(Container)
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)
    stack.Push(4)
    stack.Push(5)
    stack.Push(6)
    stack.Push(7)
    fmt.Println(stack.PopBottom())
    fmt.Println(stack.PopBottom())
    fmt.Println(stack.PopBottom())
    fmt.Println(stack.PopTop())

    fmt.Println(stack.PopTop())

    fmt.Println(stack.PopTop())

}*/
