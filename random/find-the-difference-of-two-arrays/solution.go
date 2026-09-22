package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}
	result := findDifference(nums1, nums2)
	fmt.Println(result)
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	mp1 := map[int]struct{}{}
	for _, n := range nums1 {
		mp1[n] = struct{}{}
	}

	mp2 := map[int]struct{}{}
	set2 := map[int]struct{}{}
	for _, n := range nums2 {
		mp2[n] = struct{}{}
		if _, ok1 := mp1[n]; !ok1 {
			set2[n] = struct{}{}
		}
	}

	set1 := map[int]struct{}{}
	for _, n := range nums1 {
		if _, ok := mp2[n]; !ok {
			set1[n] = struct{}{}
		}
	}

	res := make([][]int, 2)
	for k, _ := range set1 {
		res[0] = append(res[0], k)
	}
	for k, _ := range set2 {
		res[1] = append(res[1], k)
	}

	return res
}