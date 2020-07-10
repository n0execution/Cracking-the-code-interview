class GraphQueueEmptyException(Exception):
    def __init__(self):
        super().__init__(f"Graph Queue is empty")
