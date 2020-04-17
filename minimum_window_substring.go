package string


func minWindow(s string, t string) string {
	
	freq, counter := make([]int, 128), make([]int, 128)
	for _, v := range t {
		freq[v]++
		counter[v]++
	}
	l := len(t)
	minV := len(s)
	p, e :=  0,0
	cnt, start := 0, 0
	for  end, v := range s {
		if freq[v] != 0  {
			if counter[v] > 0 {
				cnt++
			}
			counter[v]--
		} 
		//full
        for start < len(s) && (cnt == l || freq[s[start]] == 0) {
			if cnt == l && end - start < minV {
				minV = end - start
				p, e = start, end+1
			}
			if freq[s[start]] !=0 {
				counter[s[start]]++  
				if counter[s[start]] > 0 {
					cnt--
				}
			} 
			start++
		}
		end++
	} 
    return s[p: e]
}