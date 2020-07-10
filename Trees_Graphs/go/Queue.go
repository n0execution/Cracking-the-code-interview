package types


import (
  "fmt"
)


type Queue struct {
  elements []string
  size int
}


func (q *Queue) Enqueue(str string) {
  q.elements = append(q.elements, str)
  q.size++
}


func (q *Queue) IsEmpty() bool {
  return q.size == 0
}


func (q *Queue) Dequeue() (string, bool) {
  if q.IsEmpty() {
    return "", false
  }

  var first string
  first = q.elements[0]
  q.elements = q.elements[1:]
  q.size--

  if q.IsEmpty() {
    return first, false
  }

  return first, true
}


func (q *Queue) PrintQueue() {
  for i := 0; i < q.size; i++ {
    fmt.Printf(q.elements[i])
    fmt.Printf(", ")
  }
  fmt.Printf("\n")
}
