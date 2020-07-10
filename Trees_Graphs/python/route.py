from GraphQueue import GraphQueue
from Graph import GraphNode


def has_route(node1, node2):
    queue = GraphQueue()
    node1.marked = True
    queue.enqueue(node1)

    while not queue.is_empty():
        node = queue.dequeue()

        for n in node.neighbours:
            if n == node2:
                return True
            elif not n.marked:
                n.marked = True
                queue.enqueue(n)

    return False


def main():
    node_a = GraphNode("A")
    node_b = GraphNode("B")
    node_c = GraphNode("C")
    node_d = GraphNode("D")
    node_e = GraphNode("E")
    node_f = GraphNode("F")

    node_a.update_neighbours([node_b, node_e, node_f])

    node_b.update_neighbours([node_e])

    node_c.update_neighbours([node_b])

    node_d.update_neighbours([node_c, node_e])

    node_f.update_neighbours([node_b])

    print(has_route(node_f, node_e))


if __name__ == '__main__':
    main()
