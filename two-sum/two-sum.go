package twosum

func twoSum(nums []int, target int) []int {
	memo := make(map[int]int)
	for i, v := range nums {
		remainder := target - v

		otherIndex, found := memo[remainder]
		if found {
			return []int{otherIndex, i}
		}
		memo[v] = i
	}

	return []int{-1, -1}
}
