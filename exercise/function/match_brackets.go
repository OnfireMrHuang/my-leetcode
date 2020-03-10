package function

import "github.com/golang-collections/collections/stack"


func BracketIsValid(str string)  bool {

	s := stack.New()
	mp := map[string]string{"(":")","{":"}","[":"]"}

	for i := 0; i < len(str); i++ {
		c := string(str[i])
		if _,ok := mp[c];ok{
			s.Push(c)
		}
		tmp := s.Pop().(string)
		if mp[tmp] != c {
			return false
		}
	}
	if s.Len() != 0 {
		return false
	}
	return true
}