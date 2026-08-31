class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        res=""
        for index,value in enumerate(strs[0]):
            for word in strs[1:]:
                if len(word) == index or word[index]!=value:
                    return strs[0][:index]
          
        return strs[0]
                