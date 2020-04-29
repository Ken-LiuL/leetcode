package array


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
//greedy
//evaluate how far we can move, from the start
func canJump(nums []int) bool {
   max_possible_moves := 0 
   for pos, moves := range nums {
	   if pos > max_possible_moves {
		   return false
	   }
	   //we update the maxximum possible moves from 0
	   //if we first move to pos and then find pos + moves is larger
	   //then we should update the maximum possible moves
	   max_possible_moves = max(max_possible_moves, pos + moves)
   }
   return true
}