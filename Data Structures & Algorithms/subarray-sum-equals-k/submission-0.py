class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        prefixSum={0:1} # sum:count empty subarray is of 1 count assume
        res=0
        cur_sum=0
        for n in nums:
            cur_sum+=n
            diff=cur_sum-k
            res+=prefixSum.get(diff,0)
            prefixSum[cur_sum]=prefixSum.get(cur_sum,0) + 1
        return res

