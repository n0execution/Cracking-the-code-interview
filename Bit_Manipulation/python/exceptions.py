class NoOnesException(Exception):
    def __init__(self, num):
        super().__init__(f"No ones in {num}")
