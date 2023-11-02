package one

import (
	"fmt"
	"math"
)

func equation(a float64, b float64, c float64) bool {
	return (a+b)+c == a+(b+c)
}

func printAssoziativ() {
	fmt.Println("-------------------------------------------")
	fmt.Println("4 Assoziativ")

	var a = math.Pow(10, 25)
	var b = math.Pow(-10, 25)
	var c = 1.0

	fmt.Println("Die Gleichung ist in go folgendes:", equation(a, b, c))
	fmt.Println("-------------------------------------------")
}
