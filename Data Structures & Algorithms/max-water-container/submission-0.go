func maxArea(heights []int) int {
    start:=0
    end:=len(heights)-1
    result:=0
    for start<end {
        height:=min(heights[start],heights[end])
        breadth:=end-start
        area:=height*breadth
        result=max(result,area)
        if heights[start] < heights[end] {
            start++
        } else{
            end--
        }
    }
    return result

}
