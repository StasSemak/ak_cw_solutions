package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return x * math.Sin(x)
}

func printXs(xs []float64) {
	fmt.Printf("x:\t")
	for _, x := range xs {
		fmt.Printf("%.2f\t", x)
	}
	fmt.Println()
}
func printFxs(fxs []float64) {
	fmt.Printf("f(x):\t")
	for _, fx := range fxs {
		fmt.Printf("%.6f\t", fx)
	}
	fmt.Println()
}

func main() {
	var a, b, m float64 = 1, 2.5, 30
	n := m
	h := (b - a) / n

	xs := []float64{}
	fxs := []float64{}

	var sum float64
	for i := 1; i <= int(n); i++ {
		x := a + float64(i)*h
		fx := f(x)

		xs = append(xs, x)
		fxs = append(fxs, fx)
		sum += fx
	}

	printXs(xs)
	printFxs(fxs)

	fmt.Printf("Ʃf(x):\t%v\n", sum)

	u := h * sum
	fmt.Printf("∫f(x):\t%v\n", u)
}
