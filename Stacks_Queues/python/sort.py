from Stack import Stack


class SortingStack(Stack):
    def sort(self):
        temp_stack = Stack()

        while not self.is_empty():
            top = self.pop()
            temp = top
            nums = 0

            if temp_stack.is_empty():
                temp_stack.push(top)
                continue

            top_temp = temp_stack.peek()

            if top >= top_temp:
                temp_stack.push(top)
            else:
                while True:
                    if temp_stack.is_empty():
                        break

                    top_temp = temp_stack.peek()

                    if temp >= top_temp:
                        break

                    top = temp_stack.pop()
                    self.push(top)
                    nums += 1

                temp_stack.push(temp)

                for i in range(nums):
                    top = self.pop()
                    temp_stack.push(top)

        while not temp_stack.is_empty():
            top = temp_stack.pop()
            self.push(top)


def main():
    s = SortingStack()
    s.push(1)
    s.push(-23)
    s.push(0)
    s.push(3)
    s.push(2)
    s.push(10)

    print(s)
    s.sort()
    print(s)

if __name__ == '__main__':
    main()
