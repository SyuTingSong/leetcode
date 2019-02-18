package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var l1, l2 = len(nums1), len(nums2)
	var stop = (l1 + l2 + 1) / 2
	var p1, p2 = 0, 0
	var m1 = 0
	for i := 0; i < stop; i++ {
		if p1 < l1 && p2 < l2 {
			if nums1[p1] < nums2[p2] {
				m1 = nums1[p1]
				p1++
			} else {
				m1 = nums2[p2]
				p2++
			}
		} else if p1 < l1 {
			m1 = nums1[p1]
			p1++
		} else if p2 < l2 {
			m1 = nums2[p2]
			p2++
		}
	}
	if (l1+l2)%2 != 0 {
		return float64(m1)
	}

	var m2 = 0
	if p1 < l1 && p2 < l2 {
		if nums1[p1] < nums2[p2] {
			m2 = nums1[p1]
		} else {
			m2 = nums2[p2]
		}
	} else if p1 < l1 {
		m2 = nums1[p1]
	} else if p2 < l2 {
		m2 = nums2[p2]
	}
	return float64(m1+m2) * 0.5
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}
