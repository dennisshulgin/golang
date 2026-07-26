package besttimestock

func MaxProfit(prices []int) int {
	l := 0
	result := 0

	for i := range prices {
		if prices[i]-prices[l] >= 0 {
			result = max(prices[i]-prices[l], result)
		} else {
			l = i
		}
	}

	return result
}
