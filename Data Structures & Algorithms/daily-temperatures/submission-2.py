class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        result=[0]*len(temperatures)
        stack=[] # monotonic decreasing stack of indices
        for index,value in enumerate(temperatures):
            if not stack:
                stack.append(index)
            while stack and temperatures[stack[-1]] < value :
                result[stack[-1]] = index-stack[-1]
                stack.pop()
            stack.append(index)
        return result