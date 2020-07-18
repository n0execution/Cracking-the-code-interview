package main
import (
  "types"
  "fmt"
)


func getFirstCommonAncestor(node1 *types.TreeNodeWithParent, node2 *types.TreeNodeWithParent) *types.TreeNodeWithParent {
  var path1, path2 types.StringStack
  var parent1, parent2, node *types.TreeNodeWithParent
  var way1, way2 string

  parent1 = node1
  parent2 = node2

  for parent1.Parent != nil {
    if parent1.Parent.LeftChild == parent1 {
      path1.Push("left")
    } else {
      path1.Push("right")
    }

    parent1 = parent1.Parent
  }

  for parent2.Parent != nil {
    if parent2.Parent.LeftChild == parent2 {
      path2.Push("left")
    } else {
      path2.Push("right")
    }

    parent2 = parent2.Parent
  }

  node = parent1

  for {
    way1, _ = path1.Pop()
    way2, _ = path2.Pop()

    if way1 == way2 {
      if way1 == "left" {
        node = node.LeftChild
      } else {
        node = node.RightChild
      }
    } else if node1 == node  || node2 == node {
      return node.Parent
    } else {
      break
    }

  }

  return node
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

  fmt.Println(getFirstCommonAncestor(&thirdLevelThird, &thirdLevelFourth))
}
