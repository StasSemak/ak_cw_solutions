package main

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
