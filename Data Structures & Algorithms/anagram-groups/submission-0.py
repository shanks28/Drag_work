from collections import Counter,defaultdict
class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        hash_map=defaultdict(list)
        for i in strs:
            temp="".join(sorted(i))
            hash_map[temp].append(i)
        res=[]
        for key,value in hash_map.items():
            res.append(value)
        return res
