class NoPathException(Exception):
    def __init__(self):
        super().__init__(f"No path found")
