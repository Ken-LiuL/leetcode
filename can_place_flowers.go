package array

func canPlaceFlowers(flowerbed []int, n int) bool {
	res := 0
	if len(flowerbed) == 1 {
        return flowerbed[0] == 0 || n == 0
    }
	for i := 0; i < len(flowerbed); i++ {
		cond := false
		if flowerbed[i] == 0  {
			if (i > 0 && flowerbed[i-1] != 1) && (i < len(flowerbed)-1 && flowerbed[i+1] != 1) {
				cond = true
			} else if i == 0 && flowerbed[i+1] != 1 {
				cond = true
			} else if i == len(flowerbed) - 1 && flowerbed[i-1] != 1 {
				cond = true
			}
			if cond {
				flowerbed[i] = 1
				res++
			}
		}
	}
	return res >= n 
}