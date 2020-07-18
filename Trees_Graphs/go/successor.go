package main
import (
  "types"
  "fmt"
)


func getNext(node *types.TreeNodeWithParent) *types.TreeNodeWithParent {
  if node.LeftChild != nil {
    return getNext(node.LeftChild)
  }

  return node
}


func successor(node *types.TreeNodeWithParent) *types.TreeNodeWithParent {
  if node.RightChild == nil {
    parent := node.Parent
    for parent.Value < node.Value {
      parent = parent.Parent

      if parent == nil {
        return &types.TreeNodeWithParent{}
      }
    }

    return parent
  }

  return getNext(node.RightChild)
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

  fmt.Println(successor(&thirdLevelFourth))
}
