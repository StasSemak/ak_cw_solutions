package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return x * math.Sin(x)
}

func main() {
	var a, b, m float64 = 1, 2.5, 30
	n := m
	h := (b - a) / n

	fmt.Println("x\t\tf(x)")

	var sum float64
	for i := 1; i <= int(n); i++ {
		x := a + float64(i)*h
		fx := f(x)

		fmt.Printf("%.2f\t\t%.6f\n", x, fx)
		sum += fx
	}

	fmt.Printf("Ʃf(x):\t%v\n", sum)

	u := h * sum
	fmt.Printf("∫f(x):\t%v\n", u)
}
