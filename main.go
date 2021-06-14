package main

import (
	"fmt"
	"github.com/y-mabuchi/go-programming-language/ch3/comma"
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
		fmt.Println(comma.Comma(arg))
	}
}
