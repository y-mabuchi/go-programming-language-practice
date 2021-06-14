package main

import (
	"fmt"
	"github.com/y-mabuchi/go-programming-language/ch2/lengthconv"
	"github.com/y-mabuchi/go-programming-language/ch2/tempconv"
	"github.com/y-mabuchi/go-programming-language/ch2/weightconv"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Print("Input Number: ")
		var t float64
		fmt.Scan(&t)
		convUnit(t)
		convLength(t)
		convWeight(t)
		return
	}

	for _, args := range os.Args[1:] {
		t, err := strconv.ParseFloat(args, 64)
		if err != nil {
			fmt.Fprintf(os.Stdout, "parse float: %v\n", err)
			os.Exit(1)
		}
		convUnit(t)
		convLength(t)
		convWeight(t)
	}
}

func convUnit(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%g°F = %g°C, %g°C = %g°F\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
}

func convLength(t float64) {
	f := lengthconv.Feet(t)
	m := lengthconv.Meter(t)
	fmt.Printf("%gft = %gm, %gm = %gft\n", f, lengthconv.FToM(f), m, lengthconv.MToF(m))
}

func convWeight(t float64) {
	p := weightconv.Pond(t)
	k := weightconv.Kilogram(t)
	fmt.Printf("%glb = %gkg, %gkg = %glb\n", p, weightconv.PToK(p), k, weightconv.KToP(k))
}