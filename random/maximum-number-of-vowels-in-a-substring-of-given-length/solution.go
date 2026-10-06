package main

import "fmt"

func main() {
	s := "weallloveyou"
	k := 7
	result := maxVowels(s, k)
	fmt.Println(result)
}

func maxVowels(s string, k int) int {
	vowels := map[byte]struct{}{
		'a':{},
		'e':{},
		'i':{},
		'o':{},
		'u':{},
	}

	vowSum := 0
	for i := 0; i < k; i++ {
		if _, ok := vowels[s[i]]; ok {
			vowSum++
		}
	}

    res := vowSum
	for l, r := 1, k; r < len(s); l, r = l+1, r+1 {
		if _, ok := vowels[s[l-1]]; ok {
			vowSum--
		}
		if _, ok := vowels[s[r]]; ok {
			vowSum++
		}
		res = max(res, vowSum)
        if res == k {
            return res
        }
	}

	return res
}