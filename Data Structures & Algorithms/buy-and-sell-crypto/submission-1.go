func maxProfit(prices []int) int {
	var res int = 0
	var bp int = 999
	for _,value := range prices {
		if bp > value {
			bp=value
		} else {
			res=max(res,value-bp)
		}
	}
	return res
	
}
