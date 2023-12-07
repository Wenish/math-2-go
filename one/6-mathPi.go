package one

import (
	"fmt"
	"math"
)

func printMathPi() {
	fmt.Println("-------------------------------------------")
	fmt.Println("6 math.Pi ̸= π")
	fmt.Println("math.Pi", math.Pi)

	// 1. Beginne mit einem Intervall [A, B], für welches gilt: π ∈ [A, B]
	// 2. Falls A = B → Ende!
	// 3. Aus den Grenzen des Intervalls A und B berechne ein C, so dass A < C < B
	// 4. Falls π ∈ [A, C] → neues Intervall ist [A, C]
	// 5. Falls π ∈ [C, B] → neues Intervall ist [C, B]
	// 6. Gehe zu Punkt 2

	A := 2.0
	B := 1.0

	C := 4.0
	D := 1.0

	X := A / B
	Y := C / D

	for math.Abs(X-Y) > 0.000000000000001 {
		top := A + C
		bottom := B + D
		result := top / bottom

		if math.Pi <= result {
			C = top
			D = bottom
			Y = C / D
		} else {
			A = top
			B = bottom
			X = A / B
		}
	}

	fmt.Printf("π liegt im Intervall [%.64f, %.64f]\n", X, Y)
	fmt.Println("-------------------------------------------")
}
