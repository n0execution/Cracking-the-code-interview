package main

import (
  "types"
  "fmt"
)


func getChild(array []int) (*types.TreeNode) {
  length := len(array)
  node := &types.TreeNode{Value: array[length / 2]}

  if length == 1 {
    return node
  }

  if length == 2 {
    node.LeftChild = getChild(array[:length / 2])
  } else {
    node.LeftChild, node.RightChild = getChild(array[:length / 2]), getChild(array[length / 2 + 1:])
  }

  return node
}


func createTree(array []int) *types.TreeNode {
  length := len(array)
  root := &types.TreeNode{Value: array[length / 2]}

  root.LeftChild, root.RightChild = getChild(array[:length / 2]), getChild(array[length / 2  + 1:])

  return root
}


func main() {
  array := []int{1, 2, 3, 4, 5, 6, 7}
  root := createTree(array)
  fmt.Println(root.RightChild)
  fmt.Println(root.RightChild.RightChild)
}
