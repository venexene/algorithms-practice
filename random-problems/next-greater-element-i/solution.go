package main

import "fmt"

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	res := nextGreaterElement(nums1, nums2)
	fmt.Println(res)
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
    greaters := map[int]int{}
    stack := []int{}
    for i := 0; i < len(nums2); i++ {
        for len(stack) != 0 && nums2[i] > stack[len(stack)-1] {
            greaters[stack[len(stack)-1]] = nums2[i]
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, nums2[i])
    }

    res := []int{}
    for _, n := range nums1 {
        if _, ok := greaters[n]; ok {
            res = append(res, greaters[n])
        } else {
            res = append(res, -1)
        }
    }

    return res
}