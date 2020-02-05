package sort

/**
*   Golang version is much slow than Java, major reason is the allocation of slice
*   and relevant GC
 */
func sortedSquares(originalArray []int) []int {
	size := len(originalArray)
	sortedArray := make([]int, size, size)

	sp, ep, np := 0, size-1, size-1
	for sp <= ep && originalArray[sp] < 0 {
		sv, ev := -originalArray[sp], originalArray[ep]
		if sv > ev {
			sortedArray[np], sp = sv*sv, sp+1
		} else {
			sortedArray[np], ep = ev*ev, ep-1
		}
		np--
	}
	//for remain positive numbers, there is no need for comparison
	for sp <= ep {
		ev := originalArray[ep]
		sortedArray[np], ep, np = ev*ev, ep-1, np-1
	}
	return sortedArray
}
