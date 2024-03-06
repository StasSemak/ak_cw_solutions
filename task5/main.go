package main

import (
	"errors"
	"fmt"
	"math"
)

func f(x float64) float64 {
	return math.Pow(x, 3) - 4*math.Pow(x, 2) + 2
}

func sign(x float64) float64 {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}

var ITER_LIMIT int = 5000

func specificationMethod(a, b, e float64) (res float64, err error) {
	fmt.Println("i\ta\t\tb\t\tc\t\tb - a\t\tza\tzc\tz")
	for i := range ITER_LIMIT {
		za := sign(f(a))
		c := (a + b) / 2

		fmt.Printf("%v\t%.6f\t%.6f\t%.6f\t%.6f\t%v",
			i, a, b, c, b-a, za)

		if b-a <= e {
			res = c
			fmt.Println()
			return
		}

		zc := sign(f(c))
		fmt.Printf("\t%v", zc)

		if zc == 0 {
			res = c
			fmt.Println()
			return
		}

		z := sign(zc * za)
		fmt.Printf("\t%v\n", z)

		if z == 1 {
			a = c
		} else if z == -1 {
			b = c
		} else {
			err = errors.New("invalid z value")
			return
		}
	}

	err = errors.New("exceeded iterations limit")
	return
}

func main() {
	var a1, b1, a2, b2, e float64 = 0.75, 0.9, 3.75, 3.9, 0.001

	fmt.Printf("[%v; %v]:\n", a1, b1)
	res1, err := specificationMethod(a1, b1, e)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("x1 = %.6f\n\n", res1)

	fmt.Printf("[%v; %v]:\n", a2, b2)
	res2, err := specificationMethod(a2, b2, e)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("x2 = %.6f\n", res2)
}
