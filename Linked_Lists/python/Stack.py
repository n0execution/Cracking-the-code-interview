class Stack(object):
    def __init__(self):
        self.top = None
        self.elements = []

    def push(self, x):
        self.elements.append(x)
        self.top = x

    def pop(self):
        top = self.top
        length = len(self.elements)

        if not length:
            return None

        self.elements = self.elements[:length-1]

        if length == 1:
            return top

        self.top = self.elements[length-2]

        return top
