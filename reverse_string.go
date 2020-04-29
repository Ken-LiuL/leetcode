package recursive

func reverseString(s []byte)  {
	start, end := 0, len(s) - 1
	for start < end {
		sc := s[start]
		s[start] = s[end]
		s[end] = sc
		end--
		start++
	}

}