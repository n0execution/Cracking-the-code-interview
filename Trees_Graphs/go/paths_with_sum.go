package main

import (
  "fmt"
  "types"
)


func countNums(node *types.TreeNode, sum int, value int, counter *int, passed map[*types.TreeNode]bool) {
  if node.LeftChild != nil {
    if sum + node.LeftChild.Value == value && !passed[node.LeftChild] {
      passed[node.LeftChild] = true
      *counter++
    }
    countNums(node.LeftChild, node.LeftChild.Value, value, counter, passed)
    countNums(node.LeftChild, sum + node.LeftChild.Value, value, counter, passed)
  }
  if node.RightChild != nil {
    if sum + node.RightChild.Value == value && !passed[node.RightChild]{
      passed[node.LeftChild] = true
      *counter++
    }
    countNums(node.RightChild, node.RightChild.Value, value, counter, passed)
    countNums(node.RightChild, sum + node.RightChild.Value, value, counter, passed)
  }
}


func getNums(root *types.TreeNode, value int) int {
  var counter int
  passed := make(map[*types.TreeNode]bool)

  countNums(root, root.Value, value, &counter, passed)

  return counter
}


func main() {
  array := []int{12, -6, 4, 3, 2, 3, 8, 1, 3, 1, 12, 6, -14, 1, 7}
  root := types.CreateTree(array)
  types.PrintTreeNodes(root)

  fmt.Println(getNums(root, 6))
}
