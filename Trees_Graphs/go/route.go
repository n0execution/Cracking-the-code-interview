package main

import (
  "types"
  "fmt"
)


func hasRoute(node1 *types.GraphNode, node2 *types.GraphNode) bool {
  var q types.GraphQueue
  node1.Marked = true
  q.Enqueue(node1)

  for !q.IsEmpty() {
    node, _ := q.Dequeue()

    for i := 0; i < len(node.Neighbours); i++ {
      n := node.Neighbours[i]

      if n == node2 {
        return true
      } else if !n.Marked {
        n.Marked = true
        q.Enqueue(n)
      }
    }
  }

  return false
}


func main() {
  nodeA := &types.GraphNode{Value: "A"}
  nodeB := &types.GraphNode{Value: "B"}
  nodeC := &types.GraphNode{Value: "C"}
  nodeD := &types.GraphNode{Value: "D"}
  nodeE := &types.GraphNode{Value: "E"}
  nodeF := &types.GraphNode{Value: "F"}

  nodeA.Neighbours = append(nodeA.Neighbours, nodeB)
  nodeA.Neighbours = append(nodeA.Neighbours, nodeE)
  nodeA.Neighbours = append(nodeA.Neighbours, nodeF)

  nodeB.Neighbours = append(nodeB.Neighbours, nodeE)

  nodeC.Neighbours = append(nodeC.Neighbours, nodeB)

  nodeD.Neighbours = append(nodeD.Neighbours, nodeC)
  nodeD.Neighbours = append(nodeD.Neighbours, nodeE)

  nodeF.Neighbours = append(nodeF.Neighbours, nodeB)

  fmt.Println(hasRoute(nodeA, nodeC))
}
