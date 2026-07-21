package twosum

func TwoSum(nums []int, target int) []int {
	numToIndex := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		prevIndex, exists := numToIndex[target - nums[i]]

		if exists {
			return []int{prevIndex, i}
		}

		numToIndex[nums[i]] = i
	} 

	return []int{}
}