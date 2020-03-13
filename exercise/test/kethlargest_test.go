package test

import (
	"testing"
	"exercise/function/kethlargest"
)

func TestKethLargest(t *testing.T)  {
	k := 3
	arr := []int{4,5,8,2}
	kth := kethlargest.Constructor(k,arr)

	t.Log(kth.Add(3))
	t.Log(kth.Add(5))
	t.Log(kth.Add(10))
	t.Log(kth.Add(9))
	t.Log(kth.Add(4))
}