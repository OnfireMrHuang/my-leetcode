package test

import (
	"testing"
	"exercise/function"
)

func TestSlidingWindow(t *testing.T)  {

	arr := []int{-7,-8,7,5,7,1,6,0}
	l := function.MaxSlidingWindow(arr,4)
	t.Log(l)
}