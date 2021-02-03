class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        _diff2index = {}
        for i, v in enumerate(nums):
            diff = target - v
            if diff in _diff2index:
                return [_diff2index[diff], i]
            else:
                _diff2index[v] = i