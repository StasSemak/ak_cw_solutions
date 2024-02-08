package main

import "math"

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
