package array


func flipAndInvertImage(A [][]int) [][]int {
	for _, a := range A {
		start, end := 0, len(a) - 1
		for start <= end {
			if  a[start] == a[end] {
				a[start] = 1 - a[start]
				a[end] = a[start]
			}
			start++
			end--
		}
	}
	return A
}