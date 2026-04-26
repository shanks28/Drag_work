class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        # monotonic stack problem
        result=[0] * len(temperatures)
        stack=[] # will be appending only indices
        for index,value in enumerate(temperatures):
            if not stack:
                stack.append(index)
            elif temperatures[stack[-1]] > value :
                stack.append(index)
            else:
                while stack and temperatures[stack[-1]] < value:
                    popped_index=stack.pop()
                    result[popped_index]=index-popped_index
                stack.append(index)
        return result