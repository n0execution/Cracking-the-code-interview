package main
import (
  "types"
  "fmt"
)



func getIndependent(nodes []*types.DirectedGraphNode) []*types.DirectedGraphNode {
  var independentNodes []*types.DirectedGraphNode

  for i := 0; i < len(nodes); i++ {
    if len(nodes[i].Incomings) == 0 && !nodes[i].Built {
      independentNodes = append(independentNodes, nodes[i])
    }
  }

  return independentNodes
}


func remove(slice []*types.DirectedGraphNode, index int) []*types.DirectedGraphNode {
    return append(slice[:index], slice[index+1:]...)
}



func removeFromIncomings(node1 *types.DirectedGraphNode, node2 *types.DirectedGraphNode) {
  var index int

  for i := 0;  i < len(node1.Incomings); i++ {
    if node1.Incomings[i] == node2 {
      index = i
      break
    }
  }

  node1.Incomings = remove(node1.Incomings, index)
}


func RemoveEdges(node *types.DirectedGraphNode) {
  var outgoing *types.DirectedGraphNode
  for i := 0; i < len(node.Outgoings); i++ {
    outgoing = node.Outgoings[i]

    removeFromIncomings(outgoing, node)
  }
}


func build(graph *types.DirectedGraph) ([]*types.DirectedGraphNode, bool) {
    var buildOrder []*types.DirectedGraphNode
    var independentNodes []*types.DirectedGraphNode

    for len(buildOrder) < len(graph.Nodes) {
      independentNodes = getIndependent(graph.Nodes)

      for i := 0; i < len(independentNodes); i++ {
        buildOrder = append(buildOrder, independentNodes[i])

        independentNodes[i].Built = true
        RemoveEdges(independentNodes[i])
      }

      if len(buildOrder) < len(graph.Nodes) && len(independentNodes) == 0 {
        return buildOrder, false
      }
    }

    return buildOrder, true
}


func main() {
  var graph types.DirectedGraph
  nodeA := &types.DirectedGraphNode{Value: "a"}
  nodeB := &types.DirectedGraphNode{Value: "b"}
  nodeC := &types.DirectedGraphNode{Value: "c"}
  nodeD := &types.DirectedGraphNode{Value: "d"}
  nodeE := &types.DirectedGraphNode{Value: "e"}
  nodeF := &types.DirectedGraphNode{Value: "f"}

  nodeF.AddEdge(nodeA)
  nodeF.AddEdge(nodeB)

  nodeA.AddEdge(nodeD)

  nodeB.AddEdge(nodeD)

  nodeD.AddEdge(nodeC)
  nodeC.AddEdge(nodeD)

  nodes := []*types.DirectedGraphNode{nodeA, nodeB, nodeC, nodeD, nodeE, nodeF}
  graph.AppendMany(nodes...)

  order, success := build(&graph)

  if success {
    for i := 0; i < len(order); i++ {
      fmt.Println(order[i].Value)
    }
  }
}
