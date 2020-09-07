from exceptions import StackEmptyException


class Stack(object):
    def __init__(self):
        self.top = None
        self.elements = []
        self.size = 0

    def __str__(self):
        return self.elements.__str__()

    def push(self, x):
        self.elements.append(x)
        self.top = x
        self.size += 1

    def pop(self):
        top = self.top

        if self.is_empty():
            raise StackEmptyException()

        self.elements = self.elements[:self.size-1]
        self.size -= 1

        if self.size == 0:
            return top

        self.top = self.elements[self.size-1]

        return top

    def peek(self):
        top = self.top

        if self.is_empty():
            raise StackEmptyException()

        return top

    def is_empty(self):
        return self.size == 0
