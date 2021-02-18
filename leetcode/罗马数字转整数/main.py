class Solution:
    roman2Value = {
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }

    def romanToInt(self, s: str) -> int:
        if not s:
            return 0
        preValue = self.roman2Value.get(s[0], 0)
        iValue = 0
        for i in range(1, len(s)):
            currValue = self.roman2Value.get(s[i], 0)
            if preValue >= currValue:
                iValue += preValue
            else:
                iValue -= preValue
            preValue = currValue
        iValue += self.roman2Value.get(s[-1], 0)
        return iValue

print(Solution().romanToInt("LVIII"))