package math



func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
		return "0"
	}
	//import observation is that 
	//if a number with N digit multiply a number with M digit , 
	//then the  result could with maximum digit is N+M
	res := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) -1; j >= 0; j-- {
			res[i+j+1] += int(num1[i] - '0') * int(num2[j] - '0')		
            res[i + j] += res[i + j + 1] / 10
			res[i + j + 1] %= 10
		}
	}
	strRes := make([]byte, 0, 100)
	start := 0
	for start < len(res) {
		 if res[start] != 0 {
			  break
		 }
        start++
	}
	for start < len(res) {
        strRes = append(strRes, byte('0' + res[start]))
        start++
	}
	return string(strRes)
}