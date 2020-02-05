package main

import "fmt"

func kthGrammarHelper(N int, K int) byte {
	if K == 1 {
		return 0
	} else if K == 2 {
		return 1
	}
	//find the previous digit that generate current digit
	pre := kthGrammarHelper(N-1, K/2+K&1)
	fmt.Println(pre)
	//if K % 2 == 0 , it means, current digit is the second digit of generated replacement, like 0 -> 01 and 1 -> 10
	//even
	if K&1 == 0 {
		//if K is even, then this is the second digit of generated replacement
		return pre ^ 1
	} else {
		return pre
	}
}

/*
 *  The keypoint is to find the rules:
 *  kthGrammar(N, K) = kthGrammaer(N-1, (K / 2).floor ) split as first or second
 */

func kthGrammar(N int, K int) int {
	return int(kthGrammarHelper(N, K))
}
