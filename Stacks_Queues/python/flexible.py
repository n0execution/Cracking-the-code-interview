from exceptions import  WrongStackNumException


class FlexibleStack(object):
    def __init__(self, stacks_num=1):
        self.stacks_num = stacks_num
        self.elements = []
        self.lengths = {}

    def get_stack_top_index(self, stack_num):
        index = 0
        for i in range(stack_num + 1):
            index += self.lengths[i]

        return index

    def push(self, stack_num, x):
        if not 0 <= stack_num < self.stacks_num + 1:
            raise WrongStackNumException(stack_num)

        if not self.lengths.get(stack_num):
            self.lengths[stack_num] = 0

        index = self.get_stack_top_index(stack_num)
        self.elements.insert(index, x)

        self.lengths[stack_num] += 1

    def pop(self, stack_num):
        if not 0 <= stack_num < self.stacks_num + 1:
            raise WrongStackNumException(stack_num)

        index = self.get_stack_top_index(stack_num) - 1
        top = self.elements[index]

        self.elements = self.elements[:index] + self.elements[index + 1:]
        self.lengths[stack_num] -= 1

        return top


def main():
    s = FlexibleStack()

    s.push(0, 3)
    s.push(0, 35)

    print(s.elements)

    s.push(1, 1)
    s.push(1, 2)
    s.push(1, 3)

    s.push(0, 45)

    print(s.elements)
    print(s.pop(0))
    print(s.pop(1))
    print(s.pop(0))
    print(s.pop(1))
    print(s.elements)
    print(s.pop(-1))


if __name__ == '__main__':
    main()
