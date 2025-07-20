package yandex

import "fmt"

func BalanceParentheses(n int, o int, c int, buf string) {
	if o+c == 2*n {
		fmt.Println(buf)
		return
	}

	if o < n {
		BalanceParentheses(n, o+1, c, buf+"(")
	}

	if o > c {
		BalanceParentheses(n, o, c+1, buf+")")
	}

}
