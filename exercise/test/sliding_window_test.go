package test

import (
	"testing"
	"exercise/function"
)

func TestSlidingWindow(t *testing.T)  {

	arr := []int{1,3,2,4,7}
	l := function.MaxSlidingWindow(arr,3)
	t.Log(l)
}