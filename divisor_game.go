package dynamic

var memorization []bool = make([]bool, 1001)

func divisorGame(N int) bool {
	//represent whether the guy whose move first can win
	memorization[0] = false
	memorization[1] = false

	for i := 2; i <= N; i++ {
		for x := 1; x < i; x++ {
			if i % x == 0 && !memorization[i-x] {
				memorization[i] = true
				break
			}
		}
	}
	return memorization[N]
}
