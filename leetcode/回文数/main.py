class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0:
            return False
        
        y = self.reverse(x)
        return x == y

    def reverse(self, x: int) -> int:
        y = 0
        index = 0
        while x != 0:
            i = x % 10
            x = x // 10
            index += 1
            if index == 1 and i == 0:
                continue
            y = i + y*10

        if -(1 << 31) + 1 < y < 1 << 31:
            return y
        return 0