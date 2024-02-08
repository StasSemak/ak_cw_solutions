package main

import (
	"fmt"
	"math"
)

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

func multiplyRow(a [][]float64, b []float64) []float64 {
	m := makeRowMatrix(len(a))

	for i := range m {
		for j := range m {
			m[i] += a[i][j] * b[j]
		}
	}

	return m
}

func reverse(m [][]float64) [][]float64 {
	d := det(m)
	mRev := makeMatrix(len(m), len(m[0]))

	for i := range m {
		for j := range m[0] {
			mRev[j][i] = (math.Pow(-1, float64(i+j+2)) * det(subMatrix(m, i, j))) / d
		}
	}

	return mRev
}

func subMatrix(m [][]float64, iSkip int, jSkip int) [][]float64 {
	sm := makeMatrix(len(m)-1, len(m[0])-1)
	iAct, jAct := 0, 0

	for i := range m {
		if i == iSkip {
			continue
		}

		jAct = 0
		for j := range m[0] {
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

func transpile(m [][]float64) [][]float64 {
	mt := makeMatrix(len(m), len(m))

	for i := range mt {
		for j := range mt[i] {
			mt[j][i] = m[i][j]
		}
	}

	return mt
}
func multiply(a, b [][]float64) [][]float64 {
	m := makeMatrix(len(a), len(b[0]))

	for i := range m {
		for j := range m[i] {
			for k := range a[i] {
				m[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return m
}

func reversedMatrix(a [][]float64, b []float64) []float64 {
	return multiplyRow(reverse(a), b)
}

func kramer(a [][]float64, b []float64) []float64 {
	x := makeRowMatrix(len(b))
	detA := det(a)

	for i := range x {
		x[i] = det(replaceIWithB(a, i, b)) / detA
	}

	return x
}

func smallestSquare(a [][]float64, b []float64) []float64 {
	at := transpile(a)

	return multiplyRow(reverse(multiply(at, a)), multiplyRow(at, b))
}

func main() {
	a := [][]float64{
		{5, -4, 7},
		{19, -22, -15},
		{0, -4, -11},
	}
	b := []float64{-10, 10, 3}

	//method of reversed matrix
	xRev := reversedMatrix(a, b)
	printMatrix([][]float64{xRev})

	//kramer method
	xKram := kramer(a, b)
	printMatrix([][]float64{xKram})

	//method of smallest squares
	xSm := smallestSquare(a, b)
	printMatrix([][]float64{xSm})
}
