package main

import "fmt"

func getMediaOneArray(num []int, left, right int) float64 {
	for right-left > 1 {
		left = left + 1
		right = right - 1
	}
	fmt.Println(left, right, num[left], num[right])
	return float64(num[left]+num[right]*1.0) / 2
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	aSize := len(nums1)
	bSize := len(nums2)
	finalNum := 2

	if (aSize+bSize)%2 != 0 {
		//odd
		finalNum = 1
	}

	var aLeft, aRight, bLeft, bRight int
	aLeft = 0
	bLeft = 0
	aRight = aSize - 1
	bRight = bSize - 1

	for {
		if finalNum == (aSize+bSize) || aSize <= 0 || bSize <= 0 {
			break
		}

		if nums1[aLeft] < nums2[bLeft] {
			aLeft = aLeft + 1
			aSize = aSize - 1
		} else {
			bLeft = bLeft + 1
			bSize = bSize - 1
		}

		if nums1[aRight] > nums2[bRight] {
			aRight = aRight - 1
			aSize = aSize - 1
		} else {
			bRight = bRight - 1
			bSize = bSize - 1
		}
	}

	if aSize <= 0 {
		return getMediaOneArray(nums2, bLeft, bRight)
	} else if bSize <= 0 {
		return getMediaOneArray(nums1, aLeft, aRight)
	}
	return float64(nums1[aLeft]+nums2[bLeft]*1.0) / 2
}

func main() {
	a := []int{2, 3}
	b := []int{}
	fmt.Println(findMedianSortedArrays(a, b))
}
