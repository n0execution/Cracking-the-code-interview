package types


import (
  "fmt"
)


type GraphQueue struct {
  nodes []*GraphNode
  size int
}


func (q *GraphQueue) Enqueue(node *GraphNode) {
  q.nodes = append(q.nodes, node)
  q.size++
}


func (q *GraphQueue) IsEmpty() bool {
  return q.size == 0
}


func (q *GraphQueue) Dequeue() (*GraphNode, bool) {
  if q.IsEmpty() {
    return &GraphNode{}, false
  }

  var first GraphNode
  first = *q.nodes[0]
  q.nodes = q.nodes[1:]
  q.size--

  if q.IsEmpty() {
    return &first, false
  }

  return &first, true
}


func (q *GraphQueue) PrintGraphQueue() {
  for i := 0; i < q.size; i++ {
    fmt.Printf(q.nodes[i].Value)
    fmt.Printf(", ")
  }
  fmt.Printf("\n")
}
