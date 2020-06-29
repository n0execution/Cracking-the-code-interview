class Node(object):

    def __init__(self, data=None, next=None):
        self.data = data
        self.next = next

    def __str__(self):
        return str(self.data)


class LinkedList(object):
    def __init__(self):
        self.head = None
        self.tail = None

    def __str__(self):
        return '->'.join([str(data) for data in self])

    def __iter__(self):
        current = self.head
        while current:
            yield current
            current = current.next

    def __len__(self):
        length = 0
        for value in self:
            length += 1

        return length

    def append(self, data):
        if self.head is None:
            self.head = self.tail = Node(data)
        else:
            self.tail.next = Node(data)
            self.tail = self.tail.next

    def insert_at_index(self, index, data):
        if self.head is None:
            self.head = self.tail = Node(data)
        elif index == len(self):
            self.append(data)
        elif index > len(self):
            raise IndexError
        else:
            for i, node in enumerate(self):
                if i == index:
                    if node == self.head:
                        prev_head = self.head

                        self.head = Node(data)
                        self.head.next = prev_head

                        return
                    else:
                        new_node = Node(data)
                        new_node.next = node
                        prev.next = new_node
                        break
                prev = node

    def remove_at_index(self, index):
        if self.head is None or index > len(self):
            raise IndexError
        elif index == 0:
            self.head = self.head.next
            return

        prev = Node()

        for i, node in enumerate(self):
            if i == index:
                next_node = node.next
                prev.next = next_node

            prev_prev = prev
            prev = node

        if index == len(self):
            self.tail = prev_prev
            self.tail.next = next_node
