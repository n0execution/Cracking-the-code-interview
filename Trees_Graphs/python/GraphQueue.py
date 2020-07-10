from exceptions import GraphQueueEmptyException


class GraphQueue(object):
    def __init__(self):
        self.nodes = []
        self.size = 0

    def __str__(self):
        return ', '.join(nodes)

    def enqueue(self, node):
        self.nodes.append(node)
        self.size += 1

    def dequeue(self):
        if self.is_empty():
            raise GraphQueueEmptyException()

        first = self.nodes[0]
        self.nodes = self.nodes[1:]
        self.size -= 1

        return first

    def is_empty(self):
        return self.size == 0
