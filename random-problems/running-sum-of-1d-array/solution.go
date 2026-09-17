package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	res := runningSum(nums)
	fmt.Println(res)
}

func runningSum(nums []int) []int {
    run := make([]int, len(nums))
    sum := 0
    for i, n := range nums {
        sum += n
        run[i] = sum
    }
    return run
}