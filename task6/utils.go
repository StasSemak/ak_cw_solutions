package main

import "fmt"

func printMatrix(m [][]float64) {
	for i := range m {
		for j := range m[0] {
			fmt.Printf("%.2f\t", m[i][j])
		}
		fmt.Println()
	}
}

func makeMatrix(x, y int) [][]float64 {
	m := make([][]float64, x)
	for i := range m {
		m[i] = make([]float64, y)
	}
	return m
}
func makeRowMatrix(x int) []float64 {
	return make([]float64, x)
}

func copy(src [][]float64, dest *[][]float64) {
	for i := range src {
		for j := range src[0] {
			(*dest)[i][j] = src[i][j]
		}
	}
}

func replaceIWithB(a [][]float64, i int, b []float64) [][]float64 {
	m := makeMatrix(len(a), len(a[0]))
	copy(a, &m)

	for j := range m {
		m[j][i] = b[j]
	}

	return m
}
