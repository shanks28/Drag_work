func maxProfit(prices []int) int {
	//buy low sell high
	buy_price:=999
	total_profit:=0
	for _,value:=range prices {
		buy_price=min(buy_price,value)
		if value > buy_price {
			profit:=value-buy_price
			buy_price=value
			total_profit+=profit
		}
		
	}
	return total_profit
}
