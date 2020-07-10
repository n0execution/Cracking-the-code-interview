package types


type GraphNode struct {
  Value string
  Neighbours []*GraphNode
  Marked bool
}


type Graph struct {
  nodes []*GraphNode
}
