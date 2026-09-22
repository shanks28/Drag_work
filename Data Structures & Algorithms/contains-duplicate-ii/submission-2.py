class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        hash_map={}
        for j,value in enumerate(nums):
            if value not in hash_map:
                hash_map[value]=[j]
            else:
                equals=hash_map[value] # all equal positions
                for i in equals:
                    if abs(i-j)<=k:
                        return True
                hash_map[value].append(j)
        return False
