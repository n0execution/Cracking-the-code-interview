from Graph import *
from exceptions import NoValidBuildOrderError


def get_independent(nodes):
    return [node for node in nodes if len(node.incomings) == 0 and not node.built]


def build(graph):
    build_order = []

    while len(build_order) < len(graph.nodes):
        independent_nodes = get_independent(graph.nodes)

        for node in independent_nodes:
            build_order.append(node)
            node.built = True

            node.remove_edges()

        if len(build_order) < len(graph.nodes) and len(independent_nodes) == 0:
            raise NoValidBuildOrderError()

    return build_order


def main():
    graph = DirectedGraph()
    node_A = DirectedGraphNode("a")
    node_B = DirectedGraphNode("b")
    node_C = DirectedGraphNode("c")
    node_D = DirectedGraphNode("d")
    node_E = DirectedGraphNode("e")
    node_F = DirectedGraphNode("f")

    node_F.add_edge(node_A)
    node_F.add_edge(node_B)

    node_A.add_edge(node_D)

    node_B.add_edge(node_D)

    node_D.add_edge(node_C)

    graph.nodes = [node_A, node_B, node_C, node_D, node_E, node_F]

    print(build(graph))


if __name__ == '__main__':
    main()
