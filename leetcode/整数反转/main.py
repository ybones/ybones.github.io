class Solution:
    def reverse(self, x: int) -> int:
        t = abs(x)
        y = 0
        index = 0
        while t != 0:
            i = t % 10
            t = t // 10
            index += 1
            if index == 1 and i == 0:
                continue
            y = i + y*10

        if x < 0:
            y = -y

        if -(1 << 31) + 1 < y < 1 << 31:
            return y
        return 0


print(Solution().reverse(10))
