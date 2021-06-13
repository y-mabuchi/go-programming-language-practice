package main

import (
	"fmt"
	"github.com/y-mabuchi/go-programming-language/ch2/lengthconv"
	"os"
	"strconv"
)

func main() {
	for _, args := range os.Args[1:] {
		t, err := strconv.ParseFloat(args, 64)
		if err != nil {
			fmt.Fprintf(os.Stdout, "fm: %v\n", err)
			os.Exit(1)
		}
		f := lengthconv.Feet(t)
		m := lengthconv.Meter(t)
		fmt.Printf("%gft = %gm, %gm = %gft\n", f, lengthconv.FToM(f), m, lengthconv.MToF(m))
	}
}
