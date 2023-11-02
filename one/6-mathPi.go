package one

import (
	"fmt"
	"math"
)

func printMathPi() {
	fmt.Println("-------------------------------------------")
	fmt.Println("6 math.Pi ̸= π")
	fmt.Println("math.Pi", math.Pi)
	A := 3.0
	B := 3.2

	/* THIS IS SLOW AF
	for A != B {
		C := (A + B) / 2

		if math.Pi <= C {
			B = C
		} else {
			A = C
		}

		fmt.Println("[", A, C, "]")
	}
	*/
	fmt.Printf("π liegt im Intervall [%.64f, %.64f]\n", A, B)
	fmt.Println("-------------------------------------------")
}
