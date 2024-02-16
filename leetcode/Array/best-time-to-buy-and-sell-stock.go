func maxProfit(prices []int) int {
    buy,sell := 0, 1
    maxP := 0

    for r < len(prices) {
        if prices[buy] < prices[sell] {
            profit := prices[sell] - prices[buy]
            maxP = max(maxP, profit)
        } else {
            l = r    
        }
        r += 1
    }
    return maxP
}

// Usando TWO POINTERS
