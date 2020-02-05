package array

//from left to right, compare each element with its left  element
//from right to left, compare each element with its right  element
//then we know to return the sum of the rating array
func candy(ratings []int) int {
	 assignment := make([]int, len(ratings))
 	 for i := 1; i < len(ratings); i++ {
		 if  ratings[i] > ratings[i-1] {
			assignment[i] = assignment[i-1]+1
		 } 
	 }
	 for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && assignment[i] <= assignment[i+1] {
			assignment[i] = assignment[i+1]+1
		}
	 } 
	 sum := 0
	 for _, n := range assignment {
		 sum += n
	 }
    return sum+len(ratings)
}