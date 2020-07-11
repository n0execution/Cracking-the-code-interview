package main
import (
  "fmt"
  "types"
  "math"
)


func isBalanced(lchild *types.TreeNode, rchild *types.TreeNode) bool {
  var left, right bool
  if math.Abs(float64(lchild.Value - rchild.Value)) > 1 {
    return false
  }

  if lchild.LeftChild == nil || lchild.RightChild == nil{
    left = true
  } else {
    left = isBalanced(lchild.LeftChild, lchild.RightChild)
  }

  if rchild.LeftChild == nil || rchild.RightChild == nil{
    right = true
  } else {
    right = isBalanced(rchild.LeftChild, rchild.RightChild)
  }

  return left && right
}


func main() {
  array := []int{2, 5, 3, 4, 8, 6, 4, 7}
  root := types.CreateTree(array)
  types.PrintTreeNodes(root)

  fmt.Println(isBalanced(root.LeftChild, root.RightChild))
}
