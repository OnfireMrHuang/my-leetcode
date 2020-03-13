/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
type (
	Queue struct {
		start, end *node
		length int
	}
	node struct {
		value interface{}
		next *node
	}
)

// Create a new queue
func New() *Queue {
	return &Queue{nil,nil,0}
}
// Take the next item off the front of the queue
func (this *Queue) Dequeue() interface{} {
	if this.length == 0 {
		return nil
	}
	n := this.start
	if this.length == 1 {
		this.start = nil
		this.end = nil
	} else {
		this.start = this.start.next
	}
	this.length--
	return n.value
}
// Put an item on the end of a queue
func (this *Queue) Enqueue(value interface{}) {
	n := &node{value,nil}
	if this.length == 0 {
		this.start = n
		this.end = n		
	} else {
		this.end.next = n
		this.end = n
	}
	this.length++
}
// Return the number of items in the queue
func (this *Queue) Len() int {
	return this.length
}
// Return the first item in the queue without removing it
func (this *Queue) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.start.value
}

func maxSlidingWindow(nums []int, k int) []int {
	
	if len(nums) < 1 || k < 1{
		return []int{}
	}
	index := 0
	list := make([]int,len(nums))
	q := New()
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
// @lc code=end

