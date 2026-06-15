from collections import Counter
class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        counter=Counter(nums)
        res_dict=dict(sorted(counter.items(),key=lambda x:x[1],reverse=True))
        res=list(res_dict.keys())
        return res[:k]