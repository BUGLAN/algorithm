package algorithm

// BinarySearch 二分查找
func BinarySearch(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}

	low, high := 0, len(nums) -1
	var mid int

	for low<=high {
		mid = (low + high)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid -1
		} else {
			low = mid + 1
		}
	}

	return -1
}

