class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        res=""
        for index,value in enumerate(strs[0]):
            for word in strs[1:]:
                if index == len(word) or value!=word[index]:
                    return strs[0][:index]
        return strs[0]