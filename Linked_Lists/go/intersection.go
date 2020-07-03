package main
import (
  "fmt"
  "types"
)


func intersection(node1 *types.Node, node2 *types.Node) (*types.Node, bool) {
  var lastNode1, lastNode2 *types.Node
  var length1, length2 int = 1, 1
  head1, head2 := node1, node2

  for node1.Next != nil {
    lastNode1 = node1.Next
    node1 = node1.Next
    length1++
  }

  for node2.Next != nil {
    lastNode2 = node2.Next
    node2 = node2.Next
    length2++
  }

  if lastNode1 != lastNode2 {
    return &types.Node{}, false
  }

  node1, node2 = head1, head2

  if length1 > length2 {
    for i := 0; i < length1 - length2; i++ {
      node1 = node1.Next
    }
  } else if length1 < length2 {
    for i := 0; i < length2 - length1; i++ {
      node2 = node2.Next
    }
  }

  for node1.Next != nil {
    if node1.Next == node2.Next {
      return node1.Next, true
    }

    node1 = node1.Next
    node2 = node2.Next
  }

  return &types.Node{}, false

}


func main() {
  var intersectionList, list1, list2 types.LinkedList

  intersectionList.Append(1)
  intersectionList.Append(6)
  intersectionList.Append(9)

  list1.Append(2)
  list1.Append(3)

  node1 := list1.Head
  for node1.Next != nil {
    node1 = node1.Next
  }

  node1.Next = intersectionList.Head

  list2.Append(2)

  node2 := list2.Head
  for node2.Next != nil {
    node2 = node2.Next
  }

  node2.Next = intersectionList.Head

  types.PrintLinkedList(list1)
  types.PrintLinkedList(list2)

  fmt.Println(intersection(list1.Head, list2.Head))
}
