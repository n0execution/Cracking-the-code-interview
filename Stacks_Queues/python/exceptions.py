class WrongStackNumException(Exception):
    def __init__(self, stack_num):
        super().__init__(f"Wrong stack num given: {stack_num}")


class StackFullException(Exception):
    def __init__(self):
        super().__init__(f"Stack is full")


class StackEmptyException(Exception):
    def __init__(self):
        super().__init__(f"Stack is empty")


class SetEmptyException(Exception):
    def __init__(self):
        super().__init__(f"Set is empty")


class QueueEmptyException(Exception):
    def __init__(self):
        super().__init__(f"Queue is empty")
