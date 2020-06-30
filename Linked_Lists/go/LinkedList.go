package types

import (
  "fmt"
)


type Node struct {
  Data int
  Next *Node
}


type LinkedList struct {
  Head *Node
  Tail *Node
}


func (list *LinkedList) Append(data int) {
  newNode := Node{Data: data}
  if list.Head == nil {
    list.Head = &newNode
    list.Tail = &newNode
  } else {
    list.Tail.Next = &newNode
    list.Tail = list.Tail.Next
  }
}


func (list *LinkedList) AddToBeginning(data int) {
  newNode := Node{Data: data}
  if list.Head == nil {
    list.Head = &newNode
    list.Tail = &newNode
  } else {
    head := list.Head
    list.Head = &newNode
    newNode.Next = head
  }
}


func (list *LinkedList) DeleteLast() {
  node := list.Head
  prev := node

  for node.Next != nil {
    prev = node
    node = node.Next
  }

  prev.Next = nil
}


func PrintLinkedList(list LinkedList) {
  node := list.Head
  for node != nil {
    fmt.Print(node.Data, "->")
    node = node.Next
  }
  fmt.Print("nil\n")
}
