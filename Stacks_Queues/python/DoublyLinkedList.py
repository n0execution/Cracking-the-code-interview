class Node(object):

    def __init__(self, data=None, next=None, prev=None):
        self.data = data
        self.next = next
        self.prev = prev

    def __str__(self):
        return str(self.data)


class DoublyLinkedList(object):
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
            prev = self.tail
            self.tail.next = Node(data)
            self.tail = self.tail.next
            self.tail.prev = prev

    def delete_last(self):
        last = self.tail
        last.prev.next = None
        self.tail = last.prev
