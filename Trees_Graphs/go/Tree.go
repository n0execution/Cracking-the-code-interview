package types


import (
  "fmt"
)


type TreeNode struct {
  Value int
  LeftChild *TreeNode
  RightChild *TreeNode
}


func getChild(array []int) (*TreeNode) {
  length := len(array)
  node := &TreeNode{Value: array[length / 2]}

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


func CreateTree(array []int) *TreeNode {
  length := len(array)
  root := &TreeNode{Value: array[length / 2]}

  root.LeftChild, root.RightChild = getChild(array[:length / 2]), getChild(array[length / 2  + 1:])

  return root
}


type LinkedListTreeNode struct {
  Node *TreeNode
  Next *LinkedListTreeNode
}


type LinkedListOfTreeNodes struct{
  Head *LinkedListTreeNode
  Tail *LinkedListTreeNode
}


func (list *LinkedListOfTreeNodes) Append(node *TreeNode) {
  newNode := LinkedListTreeNode{Node: node}
  if list.Head == nil {
    list.Head = &newNode
    list.Tail = &newNode
  } else {
    list.Tail.Next = &newNode
    list.Tail = list.Tail.Next
  }
}


func PrintLinkedListOfTreeNodes(list LinkedListOfTreeNodes) {
  node := list.Head
  for node != nil {
    fmt.Print(node.Node.Value, "->")
    node = node.Next
  }
  fmt.Print("nil\n")
}


func writeDepth(node *TreeNode, dNum int, list *[]*LinkedListOfTreeNodes) {
  var depthList LinkedListOfTreeNodes
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


func PrintTreeNodes(root *TreeNode) {
  var depths []*LinkedListOfTreeNodes
  var depthList LinkedListOfTreeNodes

  depthList.Append(root)
  depths = append(depths, &depthList)

  writeDepth(root.LeftChild, 1, &depths)
  writeDepth(root.RightChild, 1, &depths)

  for i := 0; i < len(depths); i++ {
    PrintLinkedListOfTreeNodes(*depths[i])
  }
}
