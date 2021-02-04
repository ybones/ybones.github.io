# class Solution:
#     def lengthOfLongestSubstring(self, s: str) -> int:
#         if not s:
#             return 0

#         _maxLen = 0
#         sLen = len(s)
#         for i in range(sLen):
#             tStr = s[i]
#             for j in range(i+1, sLen):
#                 if s[j] in tStr:
#                     break
#                 tStr += s[j]
            
#             _maxLen = max(_maxLen, len(tStr))
#         return _maxLen

class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        """
        滑动窗口
        start = max(start，重复元素的index+1)
        """
        _maxLen = 0
        _dict = {}
        _start = 0
        for _end in range(len(s)):
            if s[_end] in _dict.keys():
                _start = max(_start, _dict[s[_end]] + 1)
            _dict[s[_end]] = _end
            _maxLen = max(_maxLen, _end-_start+1)
        return _maxLen