package function

import (
	"github.com/golang-collections/collections/queue"
)

func MaxSlidingWindow(nums []int, k int) []int {

	if len(nums) < 1 || k < 1{
		return []int{}
	}
	index := 0
	list := make([]int,len(nums))
	q := queue.New()
	q.Enqueue(nums[0])
	list[index] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > q.Peek().(int) {
			for q.Len() > 0 {
				q.Dequeue()
			}
		}
		q.Enqueue(nums[i])	
		if q.Len() > k {
			q.Dequeue()
		}

		list[index] = q.Peek().(int)
		if i >= k-1 {
			index++
		}
	}
	if index < 1 {
		index = 1
	}

	return list[0:index]
}