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


class NoAnimalsException(Exception):
    def __init__(self):
        super().__init__(f"Animal shelter is empty")


class NoDogsException(Exception):
    def __init__(self):
        super().__init__(f"No dogs in shelter")


class NoCatsException(Exception):
    def __init__(self):
        super().__init__(f"No cats in shelter")
