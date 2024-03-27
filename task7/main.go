package main

import (
	"fmt"
	"math"
)

func makeMatrix(x, y int) [][]float64 {
	m := make([][]float64, x)
	for i := range m {
		m[i] = make([]float64, y)
	}
	return m
}
func copy(src [][]float64, dest *[][]float64) {
	for i := range src {
		for j := range src[0] {
			(*dest)[i][j] = src[i][j]
		}
	}
}
func det(m [][]float64) float64 {
	return m[0][0]*m[1][1] - m[0][1]*m[1][0]
}

func subtituteColumn(m [][]float64, columnIndex int, column []float64) [][]float64 {
	ms := makeMatrix(len(m), len(m[0]))
	copy(m, &ms)

	for i := range ms {
		ms[i][columnIndex] = column[i]
	}

	return ms
}

func main() {
	xs := []float64{1.05, 1.37, 1.75, 2.01, 2.29, 2.55}
	ys := []float64{18.41, 21.23, 24.76, 27.11, 29.59, 31.98}

	var sumX, sumXSquared, sumY, sumXY float64
	n := len(xs)

	for _, x := range xs {
		sumX += x
		sumXSquared += math.Pow(x, 2)
	}
	for _, y := range ys {
		sumY += y
	}
	for i := range n {
		sumXY += xs[i] * ys[i]
	}

	xtx := [][]float64{{sumXSquared, sumX}, {sumX, float64(n)}}
	xty := []float64{sumXY, sumY}

	a := det(subtituteColumn(xtx, 0, xty)) / det(xtx)
	b := det(subtituteColumn(xtx, 1, xty)) / det(xtx)

	fmt.Print("xi:\t")
	for _, x := range xs {
		fmt.Printf("%v\t", x)
	}
	fmt.Print("\nyi:\t")
	for _, y := range ys {
		fmt.Printf("%v\t", y)
	}
	fmt.Printf("\na:\t%v\nb:\t%v\na*xi+b:\t", a, b)
	for _, x := range xs {
		yi := a*x + b
		fmt.Printf("%.4f\t", yi)
	}
	fmt.Println()
}
