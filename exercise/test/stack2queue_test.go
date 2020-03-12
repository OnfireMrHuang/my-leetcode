package test

import (
	"testing"
	"exercise/function/queue"
)

func TestStack2Queue(t *testing.T)  {

	// 加 -v 参数测试
	
	q := queue.New()
	q.Enqueue("1")
	q.Enqueue("2")
	q.Enqueue("3")

	t.Log(q.Dequeue())
	t.Log(q.Dequeue())
	t.Log(q.Dequeue())
	
}