

class MinStack:

    def __init__(self):
        """
        initialize your data structure here.
        """
        self._stack = []

    def push(self, x: int) -> None:
        if self._stack:
            _min = min(self._stack[-1][1], x)
        else:
            _min = x
        self._stack.append((x, _min))

    def pop(self) -> None:
        self._stack.pop(-1)

    def top(self) -> int:
        return self._stack[-1][0]

    def getMin(self) -> int:
        return self._stack[-1][1]



# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(x)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()