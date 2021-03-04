class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        len_nums = len(nums)
        j = 0
        for i in range(len_nums):
            if val != nums[i]:
                nums[j] = nums[i]
                j++
        
        return nums[:j]

