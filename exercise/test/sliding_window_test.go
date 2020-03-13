package test

import (
	"testing"
	"exercise/function"
)

func TestSlidingWindow(t *testing.T)  {

	arr := []int{1,-1,3,4}
	l := function.MaxSlidingWindow(arr,3)
	t.Log(l)
}