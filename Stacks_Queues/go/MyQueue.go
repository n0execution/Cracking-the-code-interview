package main
import (
  // "fmt"
  "types"
)


type MyQueue struct {
  queue *types.Stack
  buffer *types.Stack
}


func (q *MyQueue) Initialize() {
  q.queue = &types.Stack{}
  q.buffer = &types.Stack{}
}


func (q *MyQueue) Add(x int) {
  q.queue.Push(x)
}


func (q *MyQueue) Remove() (int, bool) {
  if q.IsEmpty() {
    return 0, false
  }
  var top, value int

  for !q.IsEmpty() {
    top, _ = q.queue.Pop()
    q.buffer.Push(top)
  }

  value, _ = q.buffer.Pop()

  for !q.buffer.IsEmpty() {
    top, _ = q.buffer.Pop()
    q.queue.Push(top)
  }

  return value, true
}


func (q *MyQueue) IsEmpty() bool {
  return q.queue.IsEmpty()
}

func PrintQueue(q MyQueue) {
  types.PrintStack(*q.queue)
}


func main() {
  var q MyQueue
  q.Initialize()

  q.Add(2)
  q.Add(5)
  q.Add(3)
  q.Add(1)
  PrintQueue(q)
  q.Remove()
  q.Remove()
  PrintQueue(q)
}
