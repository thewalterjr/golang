func maxProfit(prices []int) int {
    l,r := 0, 1 //left=buy righ=sell
    maxP := 0

    for r < len(prices) {
        //profitable?
        if prices[l] < prices[r] {
            profit := prices[r] - prices[l]
            maxP = max(maxP, profit)
        } else {
            l = r    
        }
        r += 1
    }
    return maxP
}

// Usando TWO POINTERS
