package hash

func groupAnagrams(strs []string) [][]string {
	res := make(map[string][]string)
	for _, s := range strs {
		 f := make([]byte,  26)
		 for _, c := range s {
			 f[c - 'a'] += 1
		 }
		 k := string(f)
		 if _, ok := res[k]; ok {
			res[k] = append(res[k], s)
		 } else {
			res[k] = []string{s}
		 }
	}
	ret := make([][]string, 0, len(strs))
	for _, v := range res {
		ret = append(ret, v)
	}
	return ret
}