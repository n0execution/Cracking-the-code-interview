from DoublyLinkedList import DoublyLinkedList


class StackMin(object):
    def __init__(self):
        self.elements = []
        self.mins = DoublyLinkedList()
        self.top = None

    def push(self, x):
        self.elements.append(x)
        self.top = x

        if self.mins.head == None:
            self.mins.append(x)
            return

        if x <= self.mins.tail.data:
            self.mins.append(x)

    def pop(self):
        top = self.top
        length = len(self.elements)

        if not length:
            return

        self.elements.pop()

        if top == self.mins.tail.data:
            self.mins.delete_last()

        if length != 1:
            self.top = self.elements[length-2]

        return top

    def min(self):
        return self.mins.tail.data

def main():
    s = StackMin()
    s.push(34)
    s.push(2)
    s.push(32)
    s.push(-3)
    s.pop()
    s.pop()

    print(s.elements)
    print(s.min())


if __name__ == '__main__':
    main()
