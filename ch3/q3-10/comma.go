package q3_10

import (
	"fmt"
)

func Comma(s string) string {
	if len(s) <= 3 {
		return s
	}

	buf := []byte(s)
	count := 0
	for i := len(buf) - 1; i >= 0; i-- {
		if count == 0 {
			fmt.Println(string(buf[i]))
			count++
			continue
		} else if count % 3 == 0 {
			fmt.Println(",")
		}
		fmt.Println(string(buf[i]))
		count++
	}
	return s
}
