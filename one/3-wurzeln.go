package one

import (
	"fmt"
	"math"
)

func printWurzeln() {
	fmt.Println("-------------------------------------------")
	fmt.Println("3 Wurzeln")

	var gleichung bool

	gleichung = math.Pow(math.Sqrt(2), 2) == 2

	fmt.Println("Die Gleichung ist in go folgendes:", gleichung)
	fmt.Println("-------------------------------------------")
}
