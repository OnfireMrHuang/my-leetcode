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

		if q.Len() >= k {
			q.Dequeue()
		}
		if q.Peek() != nil && nums[i] > q.Peek().(int) {
			for q.Len() > 0 {
				q.Dequeue()
			}
		}
		q.Enqueue(nums[i])	
		if i-k < 0 {
			index = 0
		} else {
			index = i-k+1
		}
		list[index] = q.Peek().(int)
	}

	return list[0:index+1]
}