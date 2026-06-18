from functools import lru_cache
class Solution:
    def climbStairs(self, n: int) -> int:
        @lru_cache
        def helper(i):
            if i<=0:
                return 0
            if i ==1:
                return 1
            if i==2:
                return 2
            return helper(i-1)+helper(i-2)
        return helper(n)
