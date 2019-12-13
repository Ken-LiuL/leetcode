package sort

/**
*  The significant point is that, if you try to compare and insert from 0, then, the time complexity will be o(mn),
*  but , try to insert backwards, from (m+n-1), and from largest to  smallest, then the time complexity will be o(max(m, n))
 */
func merge(nums1 []int, m int, nums2 []int, n int) {
	m, n, k := m-1, n-1, m+n-1
	for m >= 0 && n >= 0 {
		if nums2[n] > nums1[m] {
			nums1[k], n = nums2[n], n-1
		} else {
			nums1[k], m = nums1[m], m-1
		}
		k--
	}
	for n >= 0 {
		nums1[k], k, n = nums2[n], k-1, n-1
	}

}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
