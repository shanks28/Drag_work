func maxProfit(prices []int) int {
    // buy low sell high
    buy_price:=999
    max_profit:=-999
    for _,value := range prices {
        if value < buy_price {
            buy_price = value
        }
        max_profit=max(max_profit,value-buy_price)
        
    }
    return max_profit
}
