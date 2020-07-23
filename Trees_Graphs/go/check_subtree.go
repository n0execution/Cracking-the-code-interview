package main

import (
  "fmt"
  "types"
)


func isSubTree(node *types.TreeNodeWithParent, root2 *types.TreeNodeWithParent) bool {

  if node == root2 {
    return true
  }

  if node.LeftChild == nil && node.RightChild == nil {
    return false
  } else if node.LeftChild == nil {
    return isSubTree(node.RightChild, root2)
  } else if node.RightChild == nil {
    return isSubTree(node.LeftChild, root2)
  } else {
    return isSubTree(node.LeftChild, root2) || isSubTree(node.RightChild, root2)
  }

  return false
}


func main() {
  root := types.TreeNodeWithParent{Value: 4}

  secondLevelLeftChild := types.TreeNodeWithParent{Value: 2, Parent: &root}
  secondLevelRightChild := types.TreeNodeWithParent{Value: 6, Parent: &root}
  root.LeftChild = &secondLevelLeftChild
  root.RightChild = &secondLevelRightChild

  thirdLevelFirst := types.TreeNodeWithParent{Value: 1, Parent: &secondLevelLeftChild}
  thirdLevelSecond := types.TreeNodeWithParent{Value: 3, Parent: &secondLevelLeftChild}
  thirdLevelThird := types.TreeNodeWithParent{Value: 5, Parent: &secondLevelRightChild}
  thirdLevelFourth := types.TreeNodeWithParent{Value: 7, Parent: &secondLevelRightChild}

  secondLevelLeftChild.LeftChild = &thirdLevelFirst
  secondLevelLeftChild.RightChild = &thirdLevelSecond
  secondLevelRightChild.LeftChild = &thirdLevelThird
  secondLevelRightChild.RightChild = &thirdLevelFourth

  anotherTree := &secondLevelLeftChild

  fmt.Println(isSubTree(&root, anotherTree))
}
