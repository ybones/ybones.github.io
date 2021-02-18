class Solution:
    def longestCommonPrefix(self, strs) -> str:
        if not strs:
            return ""
        prefixStr = ""
        minStr = min(strs)
        maxStr = max(strs)
        for i in range(len(minStr)):
            if minStr[i] == maxStr[i]:
                prefixStr += minStr[i]
            else:
                break
        return prefixStr

# class Solution:
#     def longestCommonPrefix(self, strs: List[str]) -> str:
#         prefixStr = ""
#         for i in zip(*strs):
#             if len(set(i)) == 1:
#                 prefixStr += i[0]
#             else:
#                 break
#         return prefixStr