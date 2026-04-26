class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        hash_map={}
        l=0
        res=0
        maxf=0
        for r in range (len(s)):
            hash_map[s[r]] = 1+hash_map.get(s[r],0)
            maxf=max(maxf,hash_map[s[r]])
            while ((r-l+1)-maxf) > k:
                hash_map[s[l]]-=1
                l+=1
            res=max(res,r-l+1)
        return res