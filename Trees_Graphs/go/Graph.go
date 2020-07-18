package types


type GraphNode struct {
  Value string
  Neighbours []*GraphNode
  Marked bool
}


type DirectedGraphNode struct {
  Value string
  Incomings []*DirectedGraphNode
  Outgoings []*DirectedGraphNode
  Built bool
}


func (from *DirectedGraphNode) AddEdge(to *DirectedGraphNode) {
  from.Outgoings = append(from.Outgoings, to)
  to.Incomings = append(to.Incomings, from)
}


type DirectedGraph struct {
  Nodes []*DirectedGraphNode
}


func (graph *DirectedGraph) AppendMany(nodes ...*DirectedGraphNode) {
  graph.Nodes = append(graph.Nodes, nodes...)
}


type Graph struct {
  nodes []*GraphNode
}
