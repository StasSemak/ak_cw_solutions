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
	for range ITER_LIMIT {
		za := sign(f(a))

		c := (a + b) / 2
		if b-a <= e {
			res = c
			return
		}

		zc := sign(f(c))
		if zc == 0 {
			res = c
			return
		}

		z := sign(zc * za)
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
	var a, b, e float64 = 0, 3, 0.0001
	res, err := specificationMethod(a, b, e)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
}
