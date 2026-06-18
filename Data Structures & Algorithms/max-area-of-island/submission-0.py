class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        visited=set()
        rows=len(grid)
        cols=len(grid[0])
        ans=0
        def helper(i,j):
            if  i>=rows or i<0 or j>=cols or j<0 or(i,j) in visited or grid[i][j]==0:
                return 0
            visited.add((i,j))
            return 1+helper(i+1,j)+helper(i-1,j)+helper(i,j+1)+helper(i,j-1)
        for i in range(rows):
            for j in range(cols):
                res=0
                if grid[i][j]==1 and (i,j) not in visited:
                    res=helper(i,j)
                ans=max(res,ans)
        return ans
