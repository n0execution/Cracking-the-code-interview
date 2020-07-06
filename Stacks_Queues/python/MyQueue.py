from Stack import Stack
from exceptions import QueueEmptyException


class MyQueue(object):
    def __init__(self):
        self.queue = Stack()
        self.buffer = Stack()

    def add(self, x):
        self.queue.push(x)

    def remove(self):
        if self.is_empty():
            raise QueueEmptyException()

        counter = 0
        while not self.queue.is_empty():
            self.buffer.push(self.queue.pop())
            counter += 1

        value = self.buffer.pop()

        while not self.buffer.is_empty():
            self.queue.push(self.buffer.pop())

        return value

    def is_empty(self):
        return self.queue.is_empty()

    def __str__(self):
        return self.queue.__str__()


q = MyQueue()
q.add(3)
q.add(4)
q.add(5)
q.add(10)
print(q)
print(q.remove())
print(q.remove())
q.add(101)
print(q)
