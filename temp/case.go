package temp

import "math"

func reverse(x int) int {
	sign := 1

	if x < 0 {
		sign = -1
		x = -1 * x
	}

	res := 0
	for x > 0 {
		temp := x % 10
		res = res*10 + temp
		x /= 10
	}

	res = sign * res

	//	处理溢出问题
	if res > math.MaxInt32 || res < math.MinInt32 {
		res = 0
	}

	return res
}

func romanToInt(s string) int {

	chTable := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	number := 0
	for i := 0; i < len(s); i++ {
		temp, ok := chTable[s[i]]
		if ok != true {
			return number
		}

		if i > 0 && temp > chTable[s[i-1]] {
			number = number - 2*chTable[s[i-1]]
		}

		number += temp
	}
	return number
}

func searchMatrix(matrix [][]int, target int) bool {
	//二分查找法
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	if target < matrix[0][0] || target > matrix[m-1][n-1] {
		return false
	}

	//二分查找行
	r := 0
	i, j := 0, m-1
	for i < j {
		med := (i + j) / 2
		switch {
		case target < matrix[med][0]:
			j = med - 1
		case target > matrix[med][0]:
			if med == m-1 || target < matrix[med+1][0] {
				r = med
				break
			}
			i = med + 1
		default:
			return true
		}
	}

	//二分查找列
	i, j = 0, n-1
	for i <= j {
		med := (i + j) / 2
		switch {
		case matrix[r][med] < target:
			i = med + 1
		case target < matrix[r][med]:
			j = med - 1
		default:
			return true
		}
	}

	return matrix[r][j] == target
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	//首先将nums1拷贝出来
	temp := make([]int, m)
	copy(temp, nums1)

	i, j, k := 0, 0, 0
	for i = 0; i < m+n; i++ {

		if j >= m {
			nums1[i] = nums2[k]
			k++
			continue
		}

		if k >= n {
			nums1[i] = temp[j]
			j++
			continue
		}

		if temp[j] > nums2[k] {
			nums1[i] = nums2[k]
			k++
		} else {
			nums1[i] = temp[j]
			j++
		}
	}
}
