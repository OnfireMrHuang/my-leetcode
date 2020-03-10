package test

import (
	"testing"
	"exercise/function"
)

func TestMatchBrackets(t *testing.T)  {
	
	ok := function.BracketIsValid("[]{}()")
	if !ok {
		t.Fatal("test fail")
	} 
	t.Log("test success")
}