package q3_10

import (
	"fmt"
)

func Comma(s string) string {
	if len(s) <= 3 {
		return s
	}

	b := []byte(s)

	for i := len(b) - 1; i >= 0; i-- {
		fmt.Println(string(b[i]))
	}

	return s
}
