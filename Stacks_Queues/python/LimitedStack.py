from exceptions import StackFullException, StackEmptyException


class LimitedStack(object):
    def __init__(self, max_size):
        self.max_size = max_size
        self.size = 0
        self.top = None
        self.elements = []

    def __str__(self):
        return self.elements.__str__()

    def push(self, x):
        if self.is_full():
            raise StackFullException()
        self.elements.append(x)
        self.top = x
        self.size += 1

    def pop(self):
        top = self.top

        if self.is_empty():
            raise StackEmptyException()

        self.elements.pop()
        self.size -= 1

        if self.is_empty():
            return top

        self.top = self.elements[self.size-1]

        return top

    def is_full(self):
        return self.size == self.max_size

    def is_empty(self):
        return self.size == 0
