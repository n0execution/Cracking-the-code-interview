package main
import (
  "fmt"
  "types"
  "math"
)


func getHeight(node *types.TreeNode) (bool, float64) {
  var left, right float64
  var balancedLeft, balancedRight bool = true, true

  if node.LeftChild == nil && node.RightChild == nil {
    return true, 1.
  }

  if node.LeftChild == nil {
    left = 1.
  } else {
    balancedLeft, left = getHeight(node.LeftChild)
  }

  if node.RightChild == nil {
    right = 1.
  } else {
    balancedRight, right = getHeight(node.RightChild)
  }

  if balancedLeft && balancedRight {
    return true, math.Max(left + 1, right + 1)
  }

  return false, 0.0
}


func isBalanced(lchild *types.TreeNode, rchild *types.TreeNode) bool {
  var balancedLeft, balancedRight bool

  balancedLeft, _ = getHeight(lchild)
  balancedRight, _ = getHeight(rchild)
  if !balancedLeft || !balancedRight {
    return false
  }

  return true
}


func main() {
  array := []int{1, 2, 3, 4, 5, 6, 7, 8}
  root := types.CreateTree(array)
  types.PrintTreeNodes(root)

  fmt.Println(isBalanced(root.LeftChild, root.RightChild))
}
