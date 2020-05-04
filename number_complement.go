package bit


func findComplement(num int) int {
	start := 0
	for start < num {
		start = (start << 1) | 1
	}
	return start - num
}



