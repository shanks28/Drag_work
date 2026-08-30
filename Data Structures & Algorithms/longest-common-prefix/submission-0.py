class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        if not strs:
            return ""
        for index,value in enumerate(strs[0]):
            for s in strs[1:]:
                if index==len(s) or s[index]!=value:
                    return strs[0][:index]
        return strs[0]