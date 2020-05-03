package hash

func findRestaurant(list1 []string, list2 []string) []string {
	dict := make(map[string]int)
	cache := make(map[string]int)
	m := 100000
	for i, s := range list1 {
		dict[s] = i
	}
	for i, s := range list2 {
		if v, ok := dict[s]; ok {
			if v+i <= m {
				m = v + i
				cache[s] = m
			}
		}
	}
	res := make([]string, 0, 10)
	for k, v := range cache {
		if v == m {
			res = append(res, k)
		}
	}
	return res
}
