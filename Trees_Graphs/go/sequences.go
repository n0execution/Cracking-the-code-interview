package main
import (
    "fmt"
    "types"
)


func getSequences(node *types.TreeNode) [][]*types.TreeNode {
  var sequence []*types.TreeNode
  var listOfSequences, leftSequences, rightSequences [][]*types.TreeNode

  if node.LeftChild == nil && node.RightChild == nil {
    sequence = append(sequence, node)
    listOfSequences = append(listOfSequences, sequence)
  } else if node.LeftChild == nil {
    rightSequences = getSequences(node.RightChild)
    for i := 0; i < len(rightSequences); i++ {
      var s []*types.TreeNode
      s = append(s, node)
      sequence = rightSequences[i]

      s = append(s, sequence...)

      listOfSequences = append(listOfSequences, s)
    }
  } else if node.RightChild == nil {
    leftSequences = getSequences(node.LeftChild)
    for i := 0; i < len(leftSequences); i++ {
      var s []*types.TreeNode
      s = append(s, node)
      sequence = leftSequences[i]

      s = append(s, sequence...)

      listOfSequences = append(listOfSequences, s)
    }
  } else {
    leftSequences = getSequences(node.LeftChild)
    rightSequences = getSequences(node.RightChild)

    // fmt.Println("left")
    // for i := 0; i < len(leftSequences); i++ {
    //   for j := 0; j < len(leftSequences[i]); j++ {
    //     fmt.Printf("%v ", leftSequences[i][j].Value)
    //   }
    //   fmt.Println("\n")
    // }
    // fmt.Println("right")
    // for i := 0; i < len(rightSequences); i++ {
    //   for j := 0; j < len(rightSequences[i]); j++ {
    //     fmt.Printf("%v ", rightSequences[i][j].Value)
    //   }
    //   fmt.Println("\n")
    // }

    for i := 0; i < len(leftSequences); i++ {
      var begin []*types.TreeNode
      sequence = leftSequences[i]
      begin = append(begin, node)
      begin = append(begin, sequence...)

      for j := 0; j < len(rightSequences); j++ {
        var s []*types.TreeNode
        sequence = rightSequences[j]
        s = append(s, begin...)
        s = append(s, sequence...)
        listOfSequences = append(listOfSequences, s)
      }
    }

    for i := 0; i < len(rightSequences); i++ {
      var begin []*types.TreeNode
      sequence = rightSequences[i]
      begin = append(begin, node)
      begin = append(begin, sequence...)

      for j := 0; j < len(leftSequences); j++ {
        var s []*types.TreeNode
        sequence = leftSequences[j]
        s = append(s, begin...)
        s = append(s, sequence...)
        listOfSequences = append(listOfSequences, s)
      }
    }
  }

  return listOfSequences
}


func main() {
  array := []int{1, 2, 3, 4, 5, 6, 7}
  root := types.CreateTree(array)
  types.PrintTreeNodes(root)

  sequences := getSequences(root)

  for i := 0; i < len(sequences); i++ {
    for j := 0; j < len(sequences[i]); j++ {
      fmt.Printf("%v ", sequences[i][j].Value)
    }
    fmt.Println("\n")
  }
}
