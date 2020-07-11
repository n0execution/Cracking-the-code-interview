package main
import (
  "types"
)


func writeDepth(node *types.TreeNode, dNum int, list *[]*types.LinkedListOfTreeNodes) {
  var depthList types.LinkedListOfTreeNodes
  if dNum == len(*list) {
    depthList.Append(node)
    *list = append(*list, &depthList)
  } else {
    depthList = *(*list)[dNum]
    depthList.Append(node)
    (*list)[dNum] = &depthList
  }

  if node.LeftChild != nil {
    writeDepth(node.LeftChild, dNum + 1, list)
  }

  if node.RightChild != nil {
    writeDepth(node.RightChild, dNum + 1, list)
  }
}


func main() {
  array := []int{1, 2, 3, 4, 5, 6, 7, 8}
  var depths []*types.LinkedListOfTreeNodes
  var depthList types.LinkedListOfTreeNodes

  root := types.CreateTree(array)
  depthList.Append(root)
  depths = append(depths, &depthList)

  writeDepth(root.LeftChild, 1, &depths)
  writeDepth(root.RightChild, 1, &depths)

  for i := 0; i < len(depths); i++ {
    types.PrintLinkedListOfTreeNodes(*depths[i])
  }
}
