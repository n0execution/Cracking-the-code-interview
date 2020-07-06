from LimitedStack import LimitedStack
from exceptions import SetEmptyException, WrongStackNumException


class SetOfStacks(object):
    def __init__(self, max_size):
        self.max_stack_size = max_size
        self.stacks = []
        self.stacks_num = 0

    def push(self, x):
        if not self.stacks_num:
            stack = LimitedStack(self.max_stack_size)
            self.stacks.append(stack)
            self.stacks_num += 1
        else:
            stack = self.stacks[self.stacks_num - 1]

        if stack.is_full():
            stack = LimitedStack(self.max_stack_size)
            self.stacks.append(stack)
            self.stacks_num += 1

        self.stacks[self.stacks_num - 1].push(x)

    def pop(self):
        return self.pop_at(self.stacks_num - 1)


    def pop_at(self, stack_num):
        if not self.stacks_num:
            raise SetEmptyException()

        if not 0 <= stack_num < self.stacks_num:
            raise WrongStackNumException(stack_num)

        stack = self.stacks[stack_num]
        top = stack.pop()

        if stack.is_empty():
            if stack_num == self.stacks_num - 1:
                self.stacks.pop()
            else:
                self.stacks = self.stacks[:stack_num] + self.stacks[stack_num+1:]
            self.stacks_num -= 1

        return top

    def print_set_of_stacks(self):
        for stack in self.stacks:
            print(stack.elements)
        print('\n')


s = SetOfStacks(2)
s.push(2)
s.print_set_of_stacks()
s.push(4)
s.push(1)
s.push(0)
s.push(1)
s.push(10)
s.print_set_of_stacks()

print(s.pop_at(1))
print(s.pop_at(1))
s.print_set_of_stacks()
print(s.pop_at(0))
print(s.pop_at(0))
s.print_set_of_stacks()
print(s.pop())
s.print_set_of_stacks()
