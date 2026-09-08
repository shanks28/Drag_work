class Solution:
    def isValid(self, s: str) -> bool:
        stack=[] # if hash_map[closing]== top then pop
        hash_map={"]":"[","}":"{",")":"("}
        for i in s:
            if i in hash_map: # meaning closing braces
                if not stack: # invalid because no counter at all
                    return False
                elif hash_map[i]==stack[-1]:
                    stack.pop()
                else:
                    return False
            else:
                stack.append(i)
        return len(stack)==0