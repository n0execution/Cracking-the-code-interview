package main
import (
  "types"
  "fmt"
)


func validateBST(leftNode *types.TreeNode, middle *types.TreeNode, rightNode *types.TreeNode) bool {
  var left, right bool

  if leftNode == nil {
    if middle.Value > rightNode.Value {
      return false
    } else {
      return true
    }
  } else if rightNode == nil {
    if middle.Value <= leftNode.Value {
      return false
    } else {
      return true
    }
  } else {
    if leftNode.Value > middle.Value || rightNode.Value <= middle.Value {
      return false
    }
  }

  if leftNode.LeftChild == nil && leftNode.RightChild == nil {
    left = true
  } else {
    left = validateBST(leftNode.LeftChild, leftNode, leftNode.RightChild)
  }

  if rightNode.LeftChild == nil && rightNode.RightChild == nil {
    right = true
  } else {
    right = validateBST(rightNode.LeftChild, rightNode, rightNode.RightChild)
  }

  return left && right
}


func main() {
  array := []int{1, 2, 3, 4, 5, 6, 7, 8}
  root := types.CreateTree(array)
  types.PrintTreeNodes(root)

  fmt.Println(validateBST(root.LeftChild, root, root.RightChild))
}
