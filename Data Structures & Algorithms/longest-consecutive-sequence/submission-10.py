class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        nums=sorted(list(set(nums)))
        res=1
        cur_window=1
        i=1
        print(nums)
        while i < len(nums):
            if (nums[i])-(nums[i-1]) == 1:
                cur_window+=1
            else:
                cur_window=1
            i+=1
            res=max(res,cur_window)
        return res if nums else 0