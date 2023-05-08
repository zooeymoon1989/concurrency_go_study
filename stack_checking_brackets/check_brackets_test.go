package stack_checking_brackets

import "testing"

func TestCheckBrackets(t *testing.T) {
	s := "(a+(b-c))"
	result := CheckBrackets(s , 0 )
	if !result {
		t.Error("result is false")
	}
}
