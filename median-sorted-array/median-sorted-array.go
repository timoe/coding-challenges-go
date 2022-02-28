package med_sorted_array

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := merge(nums1, nums2)

	// ff length of merged array is even,
	// it needs to calculate avg. of len and len-1 index.
	if len(merged)%2 == 0 {
		n := len(merged) / 2
		return (float64(merged[n-1]) + float64(merged[n])) / 2.0
	}

	return float64(merged[len(merged)/2])
}

// merge two arrays and sort them
func merge(nums1 []int, nums2 []int) []int {
	merged := []int{}
	merged = append(merged, nums1...)
	merged = append(merged, nums2...)
	sort.Ints(merged)

	return merged
}
