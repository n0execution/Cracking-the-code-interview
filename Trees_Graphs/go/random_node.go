package main

import (
  "fmt"
  "math/rand"
)


type AnotherTreeNode struct {
  Value int
  LeftChild *AnotherTreeNode
  RightChild *AnotherTreeNode
}


type AnotherTree struct {
  Root *AnotherTreeNode
  Size int
  Index int
}


func Insert(node *AnotherTreeNode, value int) {
  if value < node.Value {
    if node.LeftChild != nil {
      Insert(node.LeftChild, value)
    } else{
      node.LeftChild = &AnotherTreeNode{Value: value}
    }
  } else if value >= node.Value{
    if node.RightChild != nil {
      Insert(node.RightChild, value)
    } else{
      node.RightChild = &AnotherTreeNode{Value: value}
    }
  }
}


func (tree *AnotherTree) InsertInOrder(value int) {
  if tree.Size == 0 {
    tree.Root = &AnotherTreeNode{Value: value}
  } else {
    Insert(tree.Root, value)
  }
  tree.Size++
}


func (tree *AnotherTree) getNode(node *AnotherTreeNode, num int) (*AnotherTreeNode) {
  var leftNode, rightNode *AnotherTreeNode

  if num == tree.Index {
    tree.Index = 0
    return node
  } else {
    tree.Index++

    if node.LeftChild == nil && node.RightChild == nil {
      return nil
    } else if node.LeftChild != nil && node.RightChild != nil {
      leftNode = tree.getNode(node.LeftChild, num)
      rightNode = tree.getNode(node.RightChild, num)

      if leftNode != nil {
        return leftNode
      } else if rightNode != nil {
        return rightNode
      }
    } else if node.LeftChild != nil {
      return tree.getNode(node.LeftChild, num)
    } else if node.RightChild != nil {
      return tree.getNode(node.RightChild, num)
    }
  }

  return nil
}


func (tree *AnotherTree) getRandomNode(root *AnotherTreeNode) (*AnotherTreeNode) {
  var itemNum int
  itemNum = rand.Intn(tree.Size)

  fmt.Println(itemNum)

  return tree.getNode(root, itemNum)
}


func (tree *AnotherTree) find(node *AnotherTreeNode, value int) (*AnotherTreeNode) {
  var leftNode, rightNode *AnotherTreeNode
  if node.Value == value {
    return node
  }

  if node.LeftChild != nil && node.RightChild != nil {
    leftNode = tree.find(node.LeftChild, value)
    rightNode = tree.find(node.RightChild, value)

    if leftNode != nil {
      return leftNode
    } else if rightNode != nil {
      return rightNode
    }
  } else if node.LeftChild != nil {
    return tree.find(node.LeftChild, value)
  } else if node.RightChild != nil {
    return tree.find(node.RightChild, value)
  }

  return nil
}


func main() {
  tree := &AnotherTree{}

  tree.InsertInOrder(5)
  tree.InsertInOrder(2)
  tree.InsertInOrder(10)
  tree.InsertInOrder(15)
  tree.InsertInOrder(11)
  tree.InsertInOrder(1)

  fmt.Println(tree.getRandomNode(tree.Root))
  fmt.Println(tree.find(tree.Root, 111))
}
