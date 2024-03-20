package main

import "fmt"

func reversedMatrix(a [][]float64, b []float64) []float64 {
	printMethodHeading("METHOD OF REVERSED MATRIX\n", a, b)

	reversed := reverse(a)
	fmt.Println("\nReversed matrix:")
	printMatrix(reversed)

	result := multiplyRow(reversed, b)
	fmt.Println("\nResult:")
	printMatrix([][]float64{result})

	return result
}

func kramer(a [][]float64, b []float64) []float64 {
	printMethodHeading("KRAMER METHOD\n", a, b)

	x := makeRowMatrix(len(b))
	detA := det(a)

	fmt.Printf("Coefficients matrix determinant: %v\n", detA)

	for i := range x {
		replaced := replaceIWithB(a, i, b)
		replacedDet := det(replaced)

		fmt.Printf("\nMatrix %v:\n", i+1)
		printMatrix(replaced)
		fmt.Printf("Matrix %v determinant: %v\n", i+1, replacedDet)

		x[i] = replacedDet / detA
	}

	fmt.Println("\nResult:")
	printMatrix([][]float64{x})

	return x
}

func smallestSquare(a [][]float64, b []float64) []float64 {
	printMethodHeading("METHOD OF SMALLEST SQUARES\n", a, b)

	at := transpile(a)
	fmt.Println("\nTranspiled coefficients matrix:")
	printMatrix(at)

	ata := multiply(at, a)
	fmt.Println("\nProduct of non- & transpiled matrixes:")
	printMatrix(ata)

	reversedAta := reverse(ata)
	fmt.Println("\nReversed product:")
	printMatrix(reversedAta)

	atb := multiplyRow(at, b)
	fmt.Println("\nProduct of transpiled matrix & free members:")
	printMatrix([][]float64{atb})

	result := multiplyRow(reversedAta, atb)
	fmt.Println("\nResult:")
	printMatrix([][]float64{result})

	return result
}

func main() {
	a := [][]float64{
		{5, -4, 7},
		{19, -22, -15},
		{0, -4, -11},
	}
	b := []float64{-10, 10, 3}

	//method of reversed matrix
	reversedMatrix(a, b)

	fmt.Println("--------------------------------")

	//kramer method
	kramer(a, b)

	fmt.Println("--------------------------------")

	//method of smallest squares
	smallestSquare(a, b)
}
