package main

import "fmt"

func main() {
    nums := []int{4,4,1,3,1,3,2,2,5,5,1,5,2,1,2,3,5,4}
    k := 2
    result := maxOperations(nums, k)
    fmt.Println(result)
}

func maxOperations(nums []int, k int) int {
    count := map[int]int{}
    for _, n := range nums {
        count[n]++
    }

    ops := 0
    for _, n := range nums {
        if count[n] > 0 {
            count[n]--
        } else {
            continue
        }

        if count[k-n] > 0 {
            count[k-n]--
            ops++
        } else {
            count[n]++
            continue
        }
    }
    return ops
}