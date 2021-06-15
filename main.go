package main

import (
	"fmt"
	q3_10 "github.com/y-mabuchi/go-programming-language/ch3/q3-10"
	"log"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		_, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalf("cannot convert to int: %v\n", err)
		}
		fmt.Println(q3_10.Comma(arg))
	}
}
