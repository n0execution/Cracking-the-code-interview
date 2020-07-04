package types

import (
  "fmt"
)


type DoublyLinkedListNode struct {
  Data int
  Next *DoublyLinkedListNode
  Prev *DoublyLinkedListNode
}


type DoublyLinkedList struct {
  Head *DoublyLinkedListNode
  Tail *DoublyLinkedListNode
}


func (list *DoublyLinkedList) Append(data int) {
  newNode := DoublyLinkedListNode{Data: data}
  if list.Head == nil {
    list.Head = &newNode
    list.Tail = &newNode
  } else {
    prev := list.Tail
    list.Tail.Next = &newNode
    list.Tail = list.Tail.Next
    list.Tail.Prev = prev
  }
}


func (list *DoublyLinkedList) AddToBeginning(data int) {
  newNode := DoublyLinkedListNode{Data: data}
  if list.Head == nil {
    list.Head = &newNode
    list.Tail = &newNode
  } else {
    head := list.Head
    list.Head = &newNode
    head.Prev = &newNode
    newNode.Next = head
  }
}


func (list *DoublyLinkedList) DeleteLast() {
  last := list.Tail
  last.Prev.Next = nil
  list.Tail = last.Prev
}


func PrintDoublyLinkedList(list DoublyLinkedList) {
  node := list.Head
  for node != nil {
    fmt.Print(node.Data, "->")
    node = node.Next
  }
  fmt.Print("nil\n")
}
