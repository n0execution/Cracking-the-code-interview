class GraphQueueEmptyException(Exception):
    def __init__(self):
        super().__init__(f"Graph Queue is empty")


class NoValidBuildOrderError(Exception):
    def __init__(self):
        super().__init__(f"No valid build order")


class StackEmptyException(Exception):
    def __init__(self):
        super().__init__(f"Stack is empty")
