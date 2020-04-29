package dp

/**
 * // This is the BinaryMatrix's API interface.
 * // You should not implement it, or speculate about its implementation
 * type BinaryMatrix struct {
 *     Get(int, int) int
 *     Dimensions() []int
 * }
 */

func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
		 dimen := binaryMatrix.Dimensions()
		 row, col := dimen[0]-1, dimen[1]-1
		 start := []int{0, col} 
		 c := -1	
		 for  start[0] <= row && start[1] >= 0 {
			if binaryMatrix.Get(start[0], start[1]) == 0 {
				start[0] += 1
			} else {
				c = start[1]
				start[1] -= 1
			} 
		 }
		 return c
}