package main

import (
	"fmt"
	"github.com/y-mabuchi/go-programming-language/ch3/basename2"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		bn := basename2.Basename(arg)
		fmt.Println(bn)
	}
}
