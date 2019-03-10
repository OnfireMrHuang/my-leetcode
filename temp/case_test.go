package temp

import (
	"fmt"
	"testing"
)

func TestRevrse(t *testing.T) {
	fmt.Println(reverse(123))
	fmt.Println(romanToInt("LVIII"))

	m := [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}
	fmt.Println(searchMatrix(m, 3))

	var nums1 = []int{1, 2, 3, 0, 0, 0, 0}
	var nums2 = []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	for i := 0; i < len(nums1); i++ {
		fmt.Println(nums1[i])
	}
}
