class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:

        rows = len(grid)
        cols = len(grid[0])

        visited = set()

        def helper(i, j):

            if (
                i < 0 or i >= rows or
                j < 0 or j >= cols or
                grid[i][j] == "0" or
                (i, j) in visited
            ):
                return

            visited.add((i, j))

            helper(i + 1, j)
            helper(i - 1, j)
            helper(i, j + 1)
            helper(i, j - 1)

        islands = 0

        for i in range(rows):
            for j in range(cols):

                if grid[i][j] == "1" and (i, j) not in visited:
                    islands += 1
                    helper(i, j)

        return islands