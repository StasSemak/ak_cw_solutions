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

	var sum float64
	for i := 1; i <= int(n); i++ {
		x := a + float64(i)*h
		sum += f(x)
	}

	u := h * sum
	fmt.Println(u)
}
