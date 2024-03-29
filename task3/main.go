package main

import (
	"fmt"
	"math"
)

func f1(x float64) float64 {
	return math.Exp(x)
}
func f2(x float64) float64 {
	return x/2 + 1/math.Pow(x, 2)
}
func f3(x float64) float64 {
	return x + math.Sin(x) - math.Cos(3*x-1)
}
func f4(x float64) float64 {
	return 1 - math.Exp(x)
}
func f5(x float64) float64 {
	return math.Exp(-(x + 7))
}

func main() {
	fns := []func(x float64) float64{f1, f2, f3, f4, f5}

	var a, b, n float64 = 10, 20, 24
	h := (b - a) / n

	fmt.Println("x\t\tF1(x)\t\t\tF2(x)\t\tF3(x)\t\tF4(x)\t\t\tF5(x)")

	for i := range int(n) {
		x := a + float64(i)*h

		fmt.Printf("%.4f\t\t", x)

		for j, f := range fns {
			if j == 4 {
				fmt.Printf("%v\t\t", f(x))
			} else {
				fmt.Printf("%.4f\t\t", f(x))
			}
		}

		fmt.Println()
	}
}
