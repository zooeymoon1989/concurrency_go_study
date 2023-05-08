package stack_checking_brackets

func CheckBrackets(s string, n int) bool {
	stack := make([]int,len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			n = n + 1
			stack[n] = 1
		}

		if s[i] == ')' {
			n = n - 1
			if n == -1 {
				break
			}
		}
	}

	if n == 0 {
		return true
	}

	return false
}
