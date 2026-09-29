package main

import (
	"fmt"
	"math"
)

func main() {
	chars := []byte{'a','a','a','b','b','a','a'}
	result := compress(chars)
	fmt.Println(result)
}

func compress(chars []byte) int {
	write := 0
	start := 0
	end := 0
	for end < len(chars) {
		for end < len(chars) && chars[end] == chars[start] {
			end++
		}
		if end - start == 1 {
			chars[write] = chars[start]
			write++
			start = end
		} else {
			chars[write] = chars[start]
			ln := end - start
			k := math.Floor(math.Log10(float64(ln)))
			i := 0
			for e := k; e >= 0; e-- {
				chars[write+i+1] = byte('0' + ln / int(math.Pow(10, e)))
				ln %= int(math.Pow(10, e))
				i++
			}
			write = write + 1 + i
			start = end
		}
	}

	return len(chars[:write])
}