class Solution:
    def maxArea(self, heights: List[int]) -> int:
        start=0
        end=len(heights)-1
        max_area=0
        while start < end :
            h=min(heights[start],heights[end])
            b=end-start
            max_area=max(max_area,b*h)
            if heights[start] > heights[end] :
                end-=1
            else:
                start+=1
        return max_area
