package queue

import "github.com/golang-collections/collections/stack"


type Queue struct {
	input *stack.Stack
	output *stack.Stack
}

func New() *Queue {
	return &Queue{
		input:stack.New(),
		output:stack.New(),
	}
}

func (q *Queue) Dequeue() interface{} {
	if q.output.Len() > 0 {
		return q.output.Pop()
	}
	for q.input.Len() > 0 {
		q.output.Push(q.input.Pop())
	}
	return q.output.Pop()
}

func (q *Queue) Enqueue(data interface{}) {
	q.input.Push(data)
}

func (q *Queue) Len() interface{} {
	return q.output.Len() + q.input.Len()
}