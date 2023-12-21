package lib

import "math"

func Find(arr []int, target int) bool {
	if len(arr) == 0 {
		return false
	}

	mid := int(math.Floor(float64(len(arr) / 2)))

	if arr[mid] == target {
		return true
	} else if arr[mid] > target {
		return Find(arr[:mid], target)
	} else {
		return Find(arr[mid+1:], target)
	}
}
