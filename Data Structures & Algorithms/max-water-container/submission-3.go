func maxArea(heights []int) int {
	if len(heights) < 2 {
		return 0
	}
	start := 0
	end := len(heights) - 1
	max_area := 0
	for start < end {
		h := min(heights[start], heights[end])
		b := end - start
		max_area = max(max_area, h*b)

		if heights[start] <= heights[end] {
			start++
		} else if heights[start] > heights[end] {
			end--
		}
	}
	return max_area
}