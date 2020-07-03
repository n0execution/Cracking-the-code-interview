from exceptions import *


class StacksInArray(object):
    def __init__(self, stacks_num, stack_size):
        self.stacks_num = stacks_num
        self.stack_size = stack_size
        self.elements = [0] * stacks_num * stack_size
        self.lengths = [0] * stacks_num

    def push(self, stack_num, x):
        if not 0 <= stack_num < self.stacks_num:
            raise WrongStackNumException(stack_num)

        if self.lengths[stack_num] == self.stack_size:
            raise StackFullException()

        start_index = stack_num * self.stack_size
        index = start_index + self.lengths[stack_num]
        self.elements[index] = x
        self.lengths[stack_num] += 1

    def pop(self, stack_num):
        if not 0 <= stack_num < self.stacks_num:
            raise WrongStackNumException(stack_num)

        if self.lengths[stack_num] == 0:
            raise StackEmptyException()

        start_index = stack_num * self.stack_size
        index = start_index + self.lengths[stack_num] - 1

        top = self.elements[index]
        self.elements[index] = 0
        self.lengths[stack_num] -= 1

        return top

def main():
    s = StacksInArray(3, 3)

    s.push(0, 3)
    s.push(0, 35)

    s.push(1, 1)
    s.push(1, 2)
    s.push(1, 3)
    s.push(1, 3)

    print(s.elements)

    print(s.pop(0))
    print(s.pop(1))

if __name__ == '__main__':
    main()
