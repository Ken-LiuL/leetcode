package hash

import (
	"math"
	"strings"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	licensePlate = strings.ToLower(licensePlate)
	minFreq := freq(licensePlate)
	minLen := math.MaxInt32
	res := ""
	for _, word := range words {
		if len(word) < minLen {
			wordFreq := freq(word)
			if largerThan(minFreq, wordFreq) {
				res = word
				minLen = len(word)
			}
		}
	}
	return res
}

func freq(str string) [26]uint8 {
	res := [26]uint8{}
	for _, i := range str {
		if i >= 'a' && i <= 'z' {
			res[i-'a']++
		}
	}
	return res
}

func largerThan(a, b [26]uint8) bool {
	for i := 0; i < len(a); i++ {
		if b[i] < a[i] {
			return false
		}
	}
	return true
}
