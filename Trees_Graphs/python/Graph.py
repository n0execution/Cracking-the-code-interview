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
    def __init__(self, nodes):
        nodes = nodes
