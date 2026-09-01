class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hash_map={}
        # diff:index
        res=[]
        for index,value in enumerate(nums):
            diff=target-value
            if value in hash_map:
                res.extend([hash_map[value],index])
                break
            hash_map[diff]=index
        return res