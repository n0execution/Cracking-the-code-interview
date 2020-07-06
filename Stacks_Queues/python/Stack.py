class Stack(object):
    def __init__(self):
        self.top = None
        self.elements = []
        self.size = 0

    def push(self, x):
        self.elements.append(x)
        self.top = x
        self.size += 1

    def pop(self):
        top = self.top
        length = len(self.elements)

        if not length:
            return None

        self.elements = self.elements[:length-1]
        self.size -= 1

        if length == 1:
            return top

        self.top = self.elements[length-2]

        return top

    def is_empty(self):
        return self.size == 0

    def __str__(self):
        return self.elements.__str__()
