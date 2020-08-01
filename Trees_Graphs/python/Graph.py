class GraphNode(object):
    def __init__(self, value):
        self.value = value
        self.neighbours = []
        self.marked = False

    def __repr__(self):
        return self.value

    def update_neighbours(self, nodes):
        self.neighbours.extend(nodes)


class Graph(object):
    def __init__(self, nodes=[]):
        nodes = nodes


class DirectedGraphNode(object):
    def __init__(self, value):
        self.value = value
        self.incomings = []
        self.outgoings = []
        self.built = False

    def __repr__(self):
        return self.value

    def __str__(self):
        return self.value

    def add_edge(self, to):
        self.outgoings.append(to)
        to.incomings.append(self)

    def remove_edges(self):
        for outgoing in self.outgoings:
            outgoing.incomings.remove(self)


class DirectedGraph(object):
    def __init__(self, nodes=[]):
        nodes = nodes

    def __repr__(self):
        return str(self.nodes)
