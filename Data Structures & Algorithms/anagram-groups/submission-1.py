class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        hash_map={}
        for index,value in enumerate(strs):
            key="".join(sorted(value))
            if key not in hash_map:
                hash_map[key]=[value]
            else:
                hash_map[key].append(value)
        res=[]
        for key,value in hash_map.items():
            res.append(value)
        return res
            

        