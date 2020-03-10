package function

import "github.com/golang-collections/collections/stack"


func BracketIsValid(str string)  bool {

	s := stack.New()
	mp := map[string]string{"(":")","{":"}","[":"]"}
	ch := map[string]bool{"(":true,")":true,"{":true,"}":true,"[":true,"]":true}

	for i := 0; i < len(str); i++ {
		c := string(str[i])
		// 过滤掉
		if _,ok := ch[c];!ok{
			continue
		}

		if _,ok := mp[c];ok{
			s.Push(c)
			continue
		}
		tmp,ok := s.Pop().(string)
		if !ok || mp[tmp] != c {
			return false
		}
	}
	if s.Len() != 0 {
		return false
	}
	return true
}