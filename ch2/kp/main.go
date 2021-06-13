package main

import (
	"fmt"
	"github.com/y-mabuchi/go-programming-language/ch2/weightconv"
	"os"
	"strconv"
)

func main() {
	for _, args := range os.Args[1:] {
		t, err := strconv.ParseFloat(args, 64)
		if err != nil {
			fmt.Fprintf(os.Stdout, "kp: %v\n", err)
			os.Exit(1)
		}
		k := weightconv.Kilogram(t)
		p := weightconv.Pond(t)
		fmt.Printf("%gkg = %glb, %glb = %gkg\n", k, weightconv.KToP(k), p, weightconv.PToK(p))
	}
}
