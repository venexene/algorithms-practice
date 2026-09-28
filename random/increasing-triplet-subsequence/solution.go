package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{0,4,2,1,0,-1,-3}
	result := increasingTriplet(nums)
	fmt.Println(result)
}

func increasingTriplet(nums []int) bool {
    first := math.MaxInt
    second := math.MaxInt
    for i := 0; i < len(nums); i++ {
        if nums[i] <= first {
            first = nums[i]
        } else if nums[i] <= second {
            second = nums[i]
        } else {
            return true
        }
    }

    return false
}