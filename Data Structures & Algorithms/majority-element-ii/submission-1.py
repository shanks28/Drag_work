class Solution:
    def majorityElement(self, nums: List[int]) -> List[int]:
        req=len(nums)//3
        counts={} # value:freq
        res=[]
        for index,value in enumerate(nums):
            counts[value]=counts.get(value,0) + 1
            if counts[value] > req and value not in res:
                res.append(value)
        return res
            