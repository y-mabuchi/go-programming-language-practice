package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start1 := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	t1 := time.Now()
	r1 := t1.Sub(start1)
	fmt.Printf("for time: %v\n", r1)

	start2 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	t2 := time.Now()
	r2 := t2.Sub(start2)
	fmt.Printf("strings time: %v\n", r2)
}
