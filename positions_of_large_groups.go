package array


func largeGroupPositions(S string) [][]int {
	  res := make([][]int, 0, 26)
	  var prev rune
	  cnt := 1
	  first := 0

	  for ind, r := range  S  {
		 if prev == r {
			cnt++
		 } else {
			if cnt >= 3 {
				res = append(res, []int{first, ind-1})
			}
			cnt = 1
			prev = r
			first = ind
		 }
	  }
	  if cnt >= 3 {
        res = append(res, []int{first, len(S)-1})
      }
	  return res
}