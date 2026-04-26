class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        hash_map={}
        l=0
        res=0
        for r in range (len(s)):
            hash_map[s[r]] = 1+hash_map.get(s[r],0)
            while ((r-l+1)-max(hash_map.values())) > k:
                hash_map[s[l]]-=1
                l+=1
            res=max(res,r-l+1)
        return res