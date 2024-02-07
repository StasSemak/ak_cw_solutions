package main

import (
	"fmt"
	"math"
)

func printMatrix(m [][]float64) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			fmt.Printf("%.2f\t", m[i][j])
		}
		fmt.Println()
	}
}

func det(m [][]float64) float64 {
	if len(m) == 3 {
		return m[0][0]*m[1][1]*m[2][2] +
			m[0][1]*m[1][2]*m[2][0] +
			m[0][2]*m[1][0]*m[2][1] -
			m[0][2]*m[1][1]*m[2][0] -
			m[0][1]*m[1][0]*m[2][2] -
			m[0][0]*m[1][2]*m[2][1]
	}
	return m[0][0]*m[1][1] - m[0][1]*m[1][0]
}

func subMatrix(m [][]float64, iSkip int, jSkip int) [][]float64 {
	sm := [][]float64{{0, 0}, {0, 0}}
	iAct, jAct := 0, 0

	for i := 0; i < len(m); i++ {
		if i == iSkip {
			continue
		}

		jAct = 0
		for j := 0; j < len(m[0]); j++ {
			if j == jSkip {
				continue
			}
			sm[iAct][jAct] = m[i][j]
			jAct++
		}

		iAct++
	}

	return sm
}
func multiplyAB(a [][]float64, b []float64) []float64 {
	m := []float64{0, 0, 0}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			m[i] += a[i][j] * b[j]
		}
	}

	return m
}
func reverse(m [][]float64) [][]float64 {
	d := det(m)
	mRev := [][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			mRev[j][i] = (math.Pow(-1, float64(i+j+2)) * det(subMatrix(m, i, j))) / d
		}
	}

	return mRev
}

func copy(src [][]float64, dest *[][]float64) {
	for i := 0; i < len(src); i++ {
		for j := 0; j < len(src[0]); j++ {
			(*dest)[i][j] = src[i][j]
		}
	}
}
func replaceIWithB(a [][]float64, i int, b []float64) [][]float64 {
	m := [][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	copy(a, &m)

	for j := 0; j < len(a[0]); j++ {
		m[j][i] = b[j]
	}

	return m
}
func kramer(a [][]float64, b []float64) []float64 {
	x := []float64{0, 0, 0}
	detA := det(a)

	for i := 0; i < len(x); i++ {
		x[i] = det(replaceIWithB(a, i, b)) / detA
	}

	return x
}

func transpile(m [][]float64) [][]float64 {
	mt := [][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			mt[j][i] = m[i][j]
		}
	}

	return mt
}
func multiply(a, b [][]float64) [][]float64 {
	m := make([][]float64, len(a))
	for i := range m {
		m[i] = make([]float64, len(b[0]))
	}

	for i := range m {
		for j := range m[i] {
			for k := range a[i] {
				m[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return m
}

func smallestSquare(a [][]float64, b []float64) []float64 {
	at := transpile(a)

	return multiplyAB(reverse(multiply(at, a)), multiplyAB(at, b))
}

func main() {
	a := [][]float64{
		{5, -4, 7},
		{19, -22, -15},
		{0, -4, -11},
	}
	b := []float64{-10, 10, 3}

	//method of reversed matrix
	xRev := multiplyAB(reverse(a), b)
	printMatrix([][]float64{xRev})

	//kramer method
	xKram := kramer(a, b)
	printMatrix([][]float64{xKram})

	//method of smallest squares
	xSm := smallestSquare(a, b)
	printMatrix([][]float64{xSm})
}
